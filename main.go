package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	_ "hutool/internal/imports"
	_ "hutool/internal/logic"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS
	device = &Device{
		VID:      "1A86",
		PID:      "7523",
		BaudRate: 115200,
	}
	server = &Server{}
)

func main() {
	if err := wails.Run(&options.App{
		Title:  server.GetName(),
		Width:  1280,
		Height: 960,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			go device.listener()
			go server.run()
		},
		Bind: []any{
			server,
		},
	}); err != nil {
		panic(err)
	}
}
