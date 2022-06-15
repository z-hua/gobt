package composite

import "github.com/z-hua/gobt"

/**
 * 序列节点
 * 若子节点返回 failure, 停止迭代, 并向父节点返回 failure
 * 若子节点返回 success, 继续执行下一子节点
 * 若所有子节点都返回 success, 向父节点返回 success
 */
type _sequence struct {
}

func NewSequence() gobt.IExecutor {
	return new(_sequence)
}

func (s *_sequence) Name() string {
	return "sequence"
}

func (s *_sequence) Enter(rt *gobt.Runtime, node *gobt.Node) {
	s.ahead(rt, node, node.Children)
}

func (s *_sequence) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Result().IsFailure() {
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

func (s *_sequence) ahead(rt *gobt.Runtime, node *gobt.Node, children []gobt.NodeId) {
	rt.SetStatus(node.Id, "children", children[1:])
	rt.Ahead(children[0], node)
}

func (s *_sequence) leave(rt *gobt.Runtime, node *gobt.Node) {
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
