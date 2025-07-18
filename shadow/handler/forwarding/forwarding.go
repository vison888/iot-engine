package forwarding

import (
	"context"

	"github.com/vison888/go-vkit/logger"
	"github.com/vison888/iot-engine/common/proto/messaging"
)

type LocalPub struct {
}

func (l *LocalPub) Publish(ctx context.Context, topic string, msg *messaging.Message) error {
	logger.Infof("LocalPub topic:%s msg:%s", topic, msg.Payload)
	return handlerNatsMsg(topic, msg)
}

func (l *LocalPub) Close() error {
	return nil
}
