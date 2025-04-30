package main

import (
	"clonecoding/internal/bootstrap"
)

func main() {
	app := bootstrap.InitApp()
	app.Run()
}
