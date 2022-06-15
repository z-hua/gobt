package condition

import "github.com/z-hua/gobt"

/**
 * 条件节点
 * callback : 回调函数
 */
type _condition struct {
	callback func(params ...interface{}) bool
}

func NewCondition(callback func(params ...interface{}) bool) gobt.IExecutor {
	return &_condition{callback: callback}
}

func (c *_condition) Name() string {
	return "condition"
}

func (c *_condition) Enter(rt *gobt.Runtime, node *gobt.Node) {
	if c.callback(rt.Params()...) {
		rt.SetSuccess()
	} else {
		rt.SetFailure()
	}
	rt.Leave(node)
}

func (c *_condition) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
