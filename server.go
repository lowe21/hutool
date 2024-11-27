package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

type Server struct{}

func (*Server) GetName() string {
	return g.Server().GetName()
}

func (*Server) GetAddress() string {
	address := g.Server().GetListenedAddress()
	if gstr.HasPrefix(address, ":") {
		address = gstr.Join([]string{"127.0.0.1", address}, "")
	}

	return address
}

func (*Server) run() {
	g.Server().Use(ghttp.MiddlewareCORS)
	g.Server().Run()
}
