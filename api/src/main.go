package main

import (
	"github.com/guipassos/go-wine-house/api/app"
	"github.com/guipassos/go-wine-house/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
