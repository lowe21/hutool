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
			server.Run()
		},
		Bind: []any{
			server,
		},
	}); err != nil {
		panic(err)
	}
}
