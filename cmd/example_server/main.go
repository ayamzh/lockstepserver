package main

import (
	"fmt"
	"github.com/byebyebruce/lockstepserver/pkg/util"
	"github.com/byebyebruce/lockstepserver/pkg/xconfig"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/byebyebruce/lockstepserver/cmd/example_server/api"
	"github.com/byebyebruce/lockstepserver/core/log4gox"
	"github.com/byebyebruce/lockstepserver/server"

	l4g "github.com/alecthomas/log4go"
)

func main() {
	//读取环境变量
	xconfig.MustInitialize()

	l4g.Close()
	l4g.AddFilter("debug logger", l4g.DEBUG, log4gox.NewColorConsoleLogWriter())

	// 启动KCP服务
	s, err := server.New(xconfig.C().GetServer())
	util.PanicIfErr(err)

	//启动HTTP服务
	_ = api.NewWebAPI(xconfig.C().GetHttpPort(), s.RoomManager())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)
	ticker := time.NewTimer(time.Minute)
	defer ticker.Stop()

	l4g.Info("[main] start...")
	// 主循环
QUIT:
	for {
		select {
		case sig := <-sigs:
			l4g.Info("Signal: %s", sig.String())
			break QUIT
		case <-ticker.C:
			// todo
			fmt.Println("room number ", s.RoomManager().RoomNum())
		}
	}
	l4g.Info("[main] quiting...")
	s.Stop()
}
