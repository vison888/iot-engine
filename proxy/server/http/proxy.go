package http

import (
	"context"
	"fmt"

	"github.com/vison888/go-vkit/logger"
	"github.com/vison888/iot-engine/common/client/nats"
	"github.com/vison888/iot-engine/proxy/app"
	"github.com/vison888/iot-engine/proxy/server"
)

func Start() {
	svcName := "Http2nats"
	pub, err := nats.NewPub(app.Cfg.Nats.Url, app.Cfg.Nats.Username, app.Cfg.Nats.Password, svcName, "")
	if err != nil {
		logger.Errorf("failed to connect to message broker: %s", err)
		return
	}
	defer pub.Close()

	ctx, cancel := context.WithCancel(context.Background())
	sc := &server.Config{
		Host: app.Cfg.Http.Address,
		Port: fmt.Sprintf("%d", app.Cfg.Http.Port),
	}

	hs := New(ctx, cancel, svcName, sc, newHandler(pub))

	hs.Start()
}
