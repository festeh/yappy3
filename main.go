package main

import (
	"embed"
	"time"

	"github.com/charmbracelet/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"yappy3/app"

)

//go:embed all:frontend/dist
var assets embed.FS


func main() {
	// Create an instance of the app structure
	mainApp := app.NewApp()
	log.SetTimeFormat(time.Stamp)
	log.SetReportCaller(true)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "yappy3",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        mainApp.Startup,

		Bind: []interface{}{
			mainApp,
			mainApp.Pomo,
			mainApp.Coach,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
