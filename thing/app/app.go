package app

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/vison888/go-vkit/grpcclient"
	"github.com/vison888/go-vkit/logger"
	"github.com/vison888/go-vkit/mysqlx"
	"github.com/vison888/go-vkit/natsx"
	shadowpb "github.com/vison888/iot-engine/shadow/proto"
	"github.com/vison888/iot-engine/thing/config"
)

type CClient struct {
	ShadowService     *shadowpb.ShadowServiceClient
	ForwardingService *shadowpb.ForwardingServiceClient
}

var (
	Cfg    *config.Config
	Mysql  *mysqlx.MysqlClient
	Nats   *natsx.NatsClient // nats客户端
	Client *CClient
)

func Init(cfgPath string) {
	initConfig(cfgPath)
	initMysql()
	initNats()
	initClient()
}

// 初始化配置， 第一步运行
func initConfig(fpath string) {
	if cfg, err := config.LoadConfig(fpath); err == nil {
		Cfg = cfg
	} else {
		panic(err)
	}
}

func initMysql() {
	c, err := mysqlx.NewClient(Cfg.Mysql.Uri, Cfg.Mysql.MaxConn, Cfg.Mysql.MaxIdel, Cfg.Mysql.MaxLifeTime)
	if err != nil {
		panic(err)
	}
	Mysql = c
}

func initNats() {
	sub := func() {
	}

	logger.Infof("nast info Url:%s, Username:%s, Password:%s", Cfg.Nats.Url, Cfg.Nats.Username, Cfg.Nats.Password)
	newOptions := make([]nats.Option, 0)
	newOptions = append(newOptions, nats.ReconnectHandler(func(c *nats.Conn) {
		logger.Infof("[nats] ReconnectHandler")
	}))

	newOptions = append(newOptions, nats.DisconnectErrHandler(func(c *nats.Conn, e error) {
		logger.Infof("[nats] DisconnectErrHandler")
	}))

	newOptions = append(newOptions, nats.ClosedHandler(func(c *nats.Conn) {
		logger.Infof("[nats] ClosedHandler")
	}))

	Nats = natsx.NewNatsClient(Cfg.Nats.Url,
		Cfg.Nats.Username, Cfg.Nats.Password, newOptions...)

	Nats.SetConnectedHandler(func(c *nats.Conn) {
		logger.Infof("[nats] nats ConnectedHandler")
		sub()
	})
	sub()
}

func initClient() {
	Client = &CClient{}
	conn1 := grpcclient.GetConnClient(
		Cfg.InnerClient.Shadow,
		grpcclient.RequestTimeout(time.Second*20),
	)
	Client.ShadowService = shadowpb.NewShadowServiceClient("shadow", conn1)

	conn2 := grpcclient.GetConnClient(
		Cfg.InnerClient.Shadow,
		grpcclient.RequestTimeout(time.Second*20),
	)
	Client.ForwardingService = shadowpb.NewForwardingServiceClient("route", conn2)
}
