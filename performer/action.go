package performer

import "github.com/z-hua/gobt"

/**
 * 动作节点
 * callback 回调函数
 */
type _action struct {
	callback func(params ...interface{}) gobt.Result
}

func NewAction(callback func(params ...interface{}) gobt.Result) gobt.IExecutor {
	return &_action{callback: callback}
}

func (a *_action) Name() string {
	return "action"
}

func (a *_action) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetResult(a.callback(rt.Params()...))
	if rt.Result().IsRunning() {
		rt.SetRunning(node.Id)
		return
	}
	rt.Leave(node)
}

func (a *_action) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
