package run

import (
	"time"

	"github.com/vison888/go-vkit/utilsx"
	"github.com/vison888/iot-engine/common/proto/messaging"
	"github.com/vison888/iot-engine/shadow/handler/forwarding"
)

type RunningContext struct {
	runingId   string
	startTime  int64
	status     int32 // 运行 成功 失败
	failReason string
	ruleInfo   *forwarding.RuleInfo
	msg        *messaging.Message
}

func NewRunningContext(ruleInfo *forwarding.RuleInfo, msg *messaging.Message) *RunningContext {
	return &RunningContext{
		runingId:  utilsx.GenUuid(),
		startTime: time.Now().UnixMilli(),
		status:    0,
		ruleInfo:  ruleInfo,
		msg:       msg,
	}
}
