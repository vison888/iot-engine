package run

import "github.com/vison888/iot-engine/common/define"

type NodeNormalNotify struct {
	NodeBase
}

func BuildNodeNormalNotify(parantNode NodeIBase, preResult *NodeExecReult, runningContext *RunningContext, runningNode *define.RuleNode) *NodeNormalNotify {
	base := BuildNodeBase(parantNode, preResult, runningContext, runningNode)
	return &NodeNormalNotify{NodeBase: *base}
}

func (n *NodeNormalNotify) Start() error {
	return nil
}
