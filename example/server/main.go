package main

import (
	"fmt"

	"github.com/walleframe/net_websocket"
	"github.com/walleframe/walle/process"
	"github.com/walleframe/walle/testpkg/wpb"
	"github.com/walleframe/walle/zaplog"
)

func main() {
	p := 12345
	fmt.Println("port:", p)

	zaplog.SetFrameLogger(zaplog.NoopLogger)
	zaplog.SetLogicLogger(zaplog.NoopLogger)
	// zaplog.SetFrameLogger(zaplog.GetLogicLogger())

	wpb.RegisterWSvcService(process.GetRouter(), &wpb.WPBSvc{})
	svc := net_websocket.NewServer()
	svc.Run(fmt.Sprintf(":%d", p))
}
