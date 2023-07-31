package network

import (
	"net"
	"sync"
)

type Server struct {
	config    *Config         // server configuration
	callback  IConnCallback   // message callbacks in connection
	protocol  IProtocol       // customize packet protocol
	exitChan  chan struct{}   // notify all goroutines to shutdown
	waitGroup *sync.WaitGroup // wait for all goroutines
	closeOnce sync.Once
	listener  net.Listener
}

// NewServer creates a server
func NewServer(config *Config, callback IConnCallback, protocol IProtocol) *Server {
	return &Server{
		config: config,
		//这里用的是LockStepServer的Callback
		callback:  callback,
		protocol:  protocol,
		exitChan:  make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}
}

type ConnectionCreator func(net.Conn, *Server) *Session

// Start starts service
func (s *Server) Start(listener net.Listener, create ConnectionCreator) {
	s.listener = listener
	s.waitGroup.Add(1)
	defer func() {
		s.waitGroup.Done()
	}()

	for {
		select {
		case <-s.exitChan:
			return

		default:
		}

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		s.waitGroup.Add(1)
		go func() {
			defer s.waitGroup.Done()
			create(conn, s).Do()
		}()
	}
}

// Stop stops service
func (s *Server) Stop() {
	s.closeOnce.Do(func() {
		close(s.exitChan)
		s.listener.Close()
	})

	s.waitGroup.Wait()
}
