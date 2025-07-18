package main

import (
	"github.com/vison888/iot-engine/thing/app"
	"github.com/vison888/iot-engine/thing/handler/product"
	"github.com/vison888/iot-engine/thing/model"
	"github.com/vison888/iot-engine/thing/server"
)

func main() {
	// 1. 初始化配置
	app.Init("./config.toml")

	model.InitTable()

	product.InitAllProduct()

	server.Start()

}
