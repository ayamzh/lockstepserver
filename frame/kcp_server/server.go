package kcp_server

import (
	"fmt"
	"github.com/byebyebruce/lockstepserver/frame/network"
	"github.com/xtaci/kcp-go"
	"net"
)

func ListenAndServe(config *network.Config, callback network.IConnCallback,
	protocol network.IProtocol) (*network.Server, error) {

	l, err := kcp.Listen(fmt.Sprintf(":%d", config.GetUdpPort()))
	if nil != err {
		return nil, err
	}

	server := network.NewServer(config, callback, protocol)
	go server.Start(l, func(conn net.Conn, i *network.Server) *network.Conn {

		// 普通模式
		// setKCPConfig(32, 32, 0, 40, 0, 0, 100, 1400)

		// 极速模式
		// setKCPConfig(32, 32, 1, 10, 2, 1, 30, 1400)

		// 普通模式：ikcp_nodelay(kcp, 0, 40, 0, 0); 极速模式： ikcp_nodelay(kcp, 1, 10, 2, 1);

		kcpConn := conn.(*kcp.UDPSession)
		kcpConn.SetNoDelay(1, 10, 2, 1)
		kcpConn.SetStreamMode(true)
		kcpConn.SetWindowSize(4096, 4096)
		kcpConn.SetReadBuffer(4 * 1024 * 1024)
		kcpConn.SetWriteBuffer(4 * 1024 * 1024)
		kcpConn.SetACKNoDelay(true)

		// 创建一个连接
		return network.NewConn(conn, server)
	})

	return server, nil
}
