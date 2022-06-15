package decorator

import "github.com/z-hua/gobt"

/**
 * 与节点
 * 相当于”逻辑与“
 */
type _and struct {
}

func NewAnd() gobt.IExecutor {
	return &_and{}
}

func (a *_and) Name() string {
	return "and"
}

func (a *_and) Enter(rt *gobt.Runtime, node *gobt.Node) {
	a.ahead(rt, node, node.Children)
}

func (a *_and) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Result().IsFailure() {
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

func (a *_and) ahead(rt *gobt.Runtime, node *gobt.Node, children []gobt.NodeId) {
	rt.SetStatus(node.Id, "children", children[1:])
	rt.Ahead(children[0], node)
}

func (a *_and) leave(rt *gobt.Runtime, node *gobt.Node) {
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
