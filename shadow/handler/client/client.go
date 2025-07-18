package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/vison888/go-vkit/grpcclient"
	"github.com/vison888/go-vkit/logger"
	"github.com/vison888/go-vkit/utilsx"
	shadowpb "github.com/vison888/iot-engine/shadow/proto"
)

func watch(cc *shadowpb.ForwardingServiceClient, pks, sns []string, msgTypes []shadowpb.MSG_TYPE, codes []string) {
	newCtx, cancelFunc := context.WithCancel(context.Background())
	defer func() {
		cancelFunc()
	}()
	aaa, err := cc.Watch(newCtx, &shadowpb.ForwardingWatchReq{
		ContextId: utilsx.GenUuid(),
		Pks:       pks,
		Sns:       sns,
		MsgTypes:  msgTypes,
		Codes:     codes,
	})

	if err != nil {
		panic(err)
	}

	logger.Infof("client end")
	for {
		resp, err := aaa.Recv()
		if err != nil {
			logger.Infof("client err:%s", err.Error())
			return
		}
		logger.Infof("client read msg:%v", resp)
	}
}

func Start() {
	conn1 := grpcclient.GetConnClient(
		"127.0.0.1:10004",
		grpcclient.RequestTimeout(time.Second*20),
	)
	cc := shadowpb.NewForwardingServiceClient("shadow", conn1)
	// go watch(cc, []string{}, []string{"sn", "sn1"}, []shadowpb.MSG_TYPE{}, []string{})

	result, err := cc.Service(context.Background(), &shadowpb.ForwardingServiceReq{
		Sn:        "sn1",
		ContextId: "123",
		Code:      "selfClean",
		Payload:   ("{\"id\":\"2\",\"context_id\":\"123\",\"time\":1524448722000,\"params\":{\"level\":1000}}"),
		Timeout:   21000,
	})

	if err != nil {
		logger.Infof("client req err:%s", err.Error())
		return
	}
	hhhh, _ := json.Marshal(result)
	logger.Infof("hhhh:%s", string(hhhh))
}
