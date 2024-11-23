package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"hutool/internal/cmd"
	_ "hutool/internal/imports"
	_ "hutool/internal/logic"
)

var assets embed.FS

func main() {
	if err := wails.Run(&options.App{
		Title:  "hutool",
		Width:  1280,
		Height: 960,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			cmd.Main.Run(ctx)
		},
	}); err != nil {
		panic(err)
	}
}
