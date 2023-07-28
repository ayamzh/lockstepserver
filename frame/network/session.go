package network

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// Error type
var (
	ErrConnClosing   = errors.New("use of closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking  = errors.New("read packet was blocking")
)

// Session exposes a set of callbacks for the various events that occur on a connection
type Session struct {
	srv               *Server
	conn              net.Conn      // the raw connection
	extraData         interface{}   // to save extra data
	closeOnce         sync.Once     // close the conn, once, per instance
	closeFlag         int32         // close flag
	closeChan         chan struct{} // close chanel
	packetSendChan    chan IPacket  // packet send chanel
	packetReceiveChan chan IPacket  // packeet receive chanel
	callback          IConnCallback // callback
}

// IConnCallback is an interface of methods that are used as callbacks on a connection
type IConnCallback interface {
	// OnConnect is called when the connection was accepted,
	// If the return value of false is closed
	OnConnect(*Session) bool

	// OnMessage is called when the connection receives a packet,
	// If the return value of false is closed
	OnMessage(*Session, IPacket) bool

	// OnClose is called when the connection closed
	OnClose(*Session)
}

// NewSession returns a wrapper of raw conn
func NewSession(conn net.Conn, srv *Server) *Session {
	return &Session{
		srv:               srv,
		callback:          srv.callback,
		conn:              conn,
		closeChan:         make(chan struct{}),
		packetSendChan:    make(chan IPacket, srv.config.PacketSendChanLimit),
		packetReceiveChan: make(chan IPacket, srv.config.PacketReceiveChanLimit),
	}
}

// GetExtraData gets the extra data from the Session
func (s *Session) GetExtraData() interface{} {
	return s.extraData
}

// PutExtraData puts the extra data with the Session
func (s *Session) PutExtraData(data interface{}) {
	s.extraData = data
}

// GetRawConn returns the raw net.TCPConn from the Session
func (s *Session) GetRawConn() net.Conn {
	return s.conn
}

// Close closes the connection
func (s *Session) Close() {
	s.closeOnce.Do(func() {
		atomic.StoreInt32(&s.closeFlag, 1)
		close(s.closeChan)
		close(s.packetSendChan)
		close(s.packetReceiveChan)
		s.conn.Close()
		s.callback.OnClose(s)
	})
}

// IsClosed indicates whether or not the connection is closed
func (s *Session) IsClosed() bool {
	return atomic.LoadInt32(&s.closeFlag) == 1
}

func (s *Session) SetCallback(callback IConnCallback) {
	s.callback = callback
}

// AsyncWritePacket async writes a packet, this method will never block
func (s *Session) AsyncWritePacket(p IPacket, timeout time.Duration) (err error) {
	if s.IsClosed() {
		return ErrConnClosing
	}

	defer func() {
		if e := recover(); e != nil {
			err = ErrConnClosing
		}
	}()

	if timeout == 0 {
		select {
		case s.packetSendChan <- p:
			return nil

		default:
			return ErrWriteBlocking
		}

	} else {
		select {
		case s.packetSendChan <- p:
			return nil

		case <-s.closeChan:
			return ErrConnClosing

		case <-time.After(timeout):
			return ErrWriteBlocking
		}
	}
}

// Do it
func (s *Session) Do() {
	if !s.callback.OnConnect(s) {
		return
	}

	asyncDo(s.handleLoop, s.srv.waitGroup)
	asyncDo(s.readLoop, s.srv.waitGroup)
	asyncDo(s.writeLoop, s.srv.waitGroup)
}

func (s *Session) readLoop() {
	defer func() {
		recover()
		s.Close()
	}()

	for {
		select {
		case <-s.srv.exitChan:
			return

		case <-s.closeChan:
			return

		default:
		}

		s.conn.SetReadDeadline(time.Now().Add(s.srv.config.ConnReadTimeout))
		p, err := s.srv.protocol.ReadPacket(s.conn)
		if err != nil {
			return
		}

		s.packetReceiveChan <- p
	}
}

func (s *Session) writeLoop() {
	defer func() {
		recover()
		s.Close()
	}()

	for {
		select {
		case <-s.srv.exitChan:
			return

		case <-s.closeChan:
			return

		case p := <-s.packetSendChan:
			if s.IsClosed() {
				return
			}
			s.conn.SetWriteDeadline(time.Now().Add(s.srv.config.ConnWriteTimeout))
			if _, err := s.conn.Write(p.Serialize()); err != nil {
				return
			}
		}
	}
}

func (s *Session) handleLoop() {
	defer func() {
		recover()
		//关闭连接、执行CALLBACK、关闭各种读写chan
		s.Close()
	}()

	for {
		select {
		//服务关闭
		case <-s.srv.exitChan:
			return
		//连接关闭
		case <-s.closeChan:
			return
		//接收数据
		case p := <-s.packetReceiveChan:
			if s.IsClosed() {
				return
			}
			//消费数据
			if !s.callback.OnMessage(s, p) {
				return
			}
		}
	}
}

func asyncDo(fn func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}
