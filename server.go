package main

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"

	"hutool/internal/pkg/websocket"
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
	go func() {
		for {
			<-time.After(time.Second)

			data := make([]map[string]int, 0, 10)
			for i := 0; i < 10; i++ {
				data = append(data, map[string]int{"x": i, "y": grand.Intn(100)})
			}
			websocket.Notice(websocket.Message("data", data))

			websocket.Notice(websocket.Message("result", map[string]int{
				"result": grand.N(-1, 1),
			}))
		}
	}()

	g.Server().Use(ghttp.MiddlewareCORS)
	g.Server().Run()
}
