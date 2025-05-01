package main

import (
	"clonecoding/internal/bootstrap"
	"clonecoding/internal/config"
)

func main() {
	config.LoadConfig(".env")

	app := bootstrap.InitApp()
	app.Run()
}
