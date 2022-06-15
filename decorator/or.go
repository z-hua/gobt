package decorator

import "github.com/z-hua/gobt"

/**
 * 或节点
 */
type _or struct {
}

func NewOr() gobt.IExecutor {
	return &_or{}
}

func (a *_or) Name() string {
	return "name"
}

func (a *_or) Enter(rt *gobt.Runtime, node *gobt.Node) {
	a.ahead(rt, node, node.Children)
}

func (a *_or) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Result().IsSuccess() {
		a.leave(rt, node)
		return
	}
	children := rt.Status(node.Id, "children").([]gobt.NodeId)
	if len(children) == 0 {
		a.leave(rt, node)
		return
	}
	a.ahead(rt, node, children)
}

func (a *_or) ahead(rt *gobt.Runtime, node *gobt.Node, children []gobt.NodeId) {
	rt.SetStatus(node.Id, "children", children[1:])
	rt.Ahead(children[0], node)
}

func (a *_or) leave(rt *gobt.Runtime, node *gobt.Node) {
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
