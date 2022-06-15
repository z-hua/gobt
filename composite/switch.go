package composite

import "github.com/z-hua/gobt"

type _switch struct {
}

func NewSwitch() gobt.IExecutor {
	return &_switch{}
}

func (s *_switch) Name() string {
	return "switch"
}

func (s *_switch) Enter(rt *gobt.Runtime, node *gobt.Node) {
	s.ahead(rt, node, node.Children)
}

func (s *_switch) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Result().IsSuccess() {
		s.leave(rt, node)
		return
	}
	children := rt.Status(node.Id, "children").([]gobt.NodeId)
	if len(children) == 0 {
		s.leave(rt, node)
		return
	}
	s.ahead(rt, node, children)
}

func (s *_switch) ahead(rt *gobt.Runtime, node *gobt.Node, children []gobt.NodeId) {
	rt.SetStatus(node.Id, "children", children[1:])
	rt.Ahead(children[0], node)
}

func (s *_switch) leave(rt *gobt.Runtime, node *gobt.Node) {
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
