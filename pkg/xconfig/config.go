package xconfig

import (
	"bytes"
	"fmt"
	"github.com/byebyebruce/lockstepserver/core/network"
	"github.com/sandwich-go/xconf"
)

const optionUsage = `
说明文档 : https://devcenter.diandian.info/docs-sandwich/quick_start/conf/conf/
`

//go:generate optiongen --option_with_struct_name=false --new_func=NewConfig --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"OptionUsage": string(optionUsage),
		"HttpPort":    80,   // @MethodComment(HTTP端口)
		"DebugLog":    true, // @MethodComment(打开DEBUGLOG)
		"Server":      (*network.Config)(network.NewServerConf()),
	}
}

type Visitor interface {
	ConfigVisitor
}

type visitor struct {
	ConfigVisitor
}

var c Visitor

func MustInitialize() {
	cc := NewConfig()
	x := xconf.New()
	err := x.Parse(cc)
	if err != nil {
		panic(err)
	}
	x.Usage()
	c = &visitor{cc}
	printConfig()
}

func printConfig() {
	//if !c.GetService().GetLogConfig() {
	//	return
	//}
	bytesBuffer := bytes.NewBuffer([]byte{})
	xconf.MustSaveVarToWriter(c, xconf.ConfigTypeYAML, bytesBuffer)
	fmt.Println(bytesBuffer)
	bytesBuffer.Reset()
}

func C() Visitor { return c }
