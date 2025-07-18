package main

import (
	"github.com/vison888/iot-engine/auth/app"
	"github.com/vison888/iot-engine/auth/model"
	"github.com/vison888/iot-engine/auth/server"
)

func main() {
	app.Init("./config.toml")
	model.InitTable()
	server.Start()
}
