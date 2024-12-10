package websocket

import (
	"context"
	"fmt"
	"sync"
)

var (
	handlers = map[string]Handler{}
	mutex    sync.Mutex
)

type Handler func(context.Context, *Client, *Input) error

// SetHandler 设置处理程序
func SetHandler(name string, handler Handler) {
	if name == "" {
		panic("websocket handler name should not be empty")
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := handlers[name]; ok {
		panic(fmt.Sprintf("duplicate websocket handler name %s", name))
	}

	handlers[name] = handler
}
