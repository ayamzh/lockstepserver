package network

import (
	"time"
)

//go:generate optiongen --option_with_struct_name=false --new_func=NewServerConf --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"UdpPort":                int(9000),                      // @MethodComment(udp端口)
		"PacketSendChanLimit":    uint32(1024),                   // @MethodComment(the limit of packet send channel)
		"PacketReceiveChanLimit": uint32(1024),                   // @MethodComment(the limit of packet receive channel)
		"ConnReadTimeout":        time.Duration(time.Second * 5), // @MethodComment(read timeout)
		"ConnWriteTimeout":       time.Duration(time.Second * 5), // @MethodComment(write timeout)
	}
}
