package server

import (
	"fmt"
	"github.com/byebyebruce/lockstepserver/frame/kcp_server"
	"github.com/byebyebruce/lockstepserver/frame/network"
	"github.com/byebyebruce/lockstepserver/frame/packet/pb_packet"
	"github.com/byebyebruce/lockstepserver/logic"
	"github.com/byebyebruce/lockstepserver/pkg/util"
)

// LockStepServer 帧同步服务器
type LockStepServer struct {
	roomMgr   *logic.RoomManager
	udpServer *network.Server
	totalConn int64
}

// New 构造帧同步服务器对象
func New(port int) (*LockStepServer, error) {
	s := &LockStepServer{
		roomMgr: logic.NewRoomManager(),
	}

	networkServer, err := kcp_server.ListenAndServe(fmt.Sprintf(":%d", port), s, &pb_packet.MsgProtocol{})
	util.PanicIfErr(err)
	s.udpServer = networkServer
	return s, nil
}

// RoomManager 获取房间管理器
func (r *LockStepServer) RoomManager() *logic.RoomManager {
	return r.roomMgr
}

// Stop 停止服务
func (r *LockStepServer) Stop() {
	r.roomMgr.Stop()
	r.udpServer.Stop()
}
