package index

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"hutool/internal/pkg/websocket"
)

func init() {
	g.Server().BindObjectRest("/", &indexController{})
}

type indexController struct{}

func (*indexController) Get(request *ghttp.Request) {
	defer func() {
		if exception := recover(); exception != nil {
			g.Log().Error(context.TODO(), exception)
		}
	}()

	websocket.Connect(request)
}
