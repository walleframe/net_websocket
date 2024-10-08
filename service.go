package net_websocket

import (
	"context"
	"net"

	"github.com/walleframe/walle/app"
)

// WsService implement app.Service interface
type WsService struct {
	svr  *WsServer
	name string
	ln   net.Listener
}

func NewService(name string, opt ...ServerOption) app.Service {
	return &WsService{
		name: name,
		svr:  NewServer(opt...),
	}
}

func (svc *WsService) Name() string {
	return svc.name
}
func (svc *WsService) Init(s app.Stoper) (err error) {
	svc.ln, err = net.Listen("tcp", svc.svr.opts.Addr)
	return
}
func (svc *WsService) Start(s app.Stoper) (err error) {
	go svc.svr.Serve(svc.ln)
	return
}
func (svc *WsService) Stop() {
	svc.svr.Shutdown(context.Background())
	return
}
func (svc *WsService) Finish() {
	return
}
