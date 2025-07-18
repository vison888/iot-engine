package main

import (
	"github.com/vison888/iot-engine/gate/app"
	"github.com/vison888/iot-engine/gate/server"
)

func main() {
	app.Init("")

	server.Start()
}
