package game

import (
	"time"

	"github.com/byebyebruce/lockstepserver/core/network"
)

type Player struct {
	id                uint64
	idx               int32
	isReady           bool //是否准备好
	isOnline          bool //是否在线
	loadingProgress   int32
	lastHeartbeatTime int64 //最后一次心跳时间
	sendFrameCount    uint32
	client            *network.Session
	info              []byte //用户信息
}

func NewPlayer(id uint64, idx int32) *Player {
	p := &Player{
		id:  id,
		idx: idx,
	}

	return p
}

func (p *Player) Connect(conn *network.Session) {
	p.client = conn
	p.isOnline = true
	p.isReady = false
	p.lastHeartbeatTime = time.Now().Unix()
}

func (p *Player) IsOnline() bool {
	return nil != p.client && p.isOnline
}

func (p *Player) RefreshHeartbeatTime() {
	p.lastHeartbeatTime = time.Now().Unix()
}

func (p *Player) GetLastHeartbeatTime() int64 {
	return p.lastHeartbeatTime
}

func (p *Player) SetSendFrameCount(c uint32) {
	p.sendFrameCount = c
}

func (p *Player) GetSendFrameCount() uint32 {
	return p.sendFrameCount
}

func (p *Player) SendMessage(msg network.IPacket) {

	if !p.IsOnline() {
		return
	}

	if nil != p.client.AsyncWritePacket(msg, 0) {
		p.client.Close()
	}
}

func (p *Player) Cleanup() {

	if nil != p.client {
		p.client.Close()
	}
	p.client = nil
	p.isReady = false
	p.isOnline = false

}
