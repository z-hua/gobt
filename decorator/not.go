package decorator

import "github.com/z-hua/gobt"

/**
 * 非节点
 * 相当于“逻辑非”
 */
type _not struct {
}

func NewNot() gobt.IExecutor {
	return &_not{}
}

func (n *_not) Name() string {
	return "not"
}

func (n *_not) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.Ahead(node.Children[0], node)
}

func (n *_not) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Result().IsSuccess() {
		rt.SetFailure()
	} else {
		rt.SetSuccess()
	}
	rt.Leave(node)
}
