package composite

import "github.com/z-hua/gobt"

/**
 * 选择节点
 * 若子节点返回 success, 停止迭代, 并向父节点返回 success
 * 若子节点返回 failure, 继续执行下一子节点
 * 若所有子节点都返回 failure, 向父节点返回 failure
 */
type _selector struct {
}

func NewSelector() gobt.IExecutor {
	return new(_selector)
}

func (s *_selector) Name() string {
	return "selector"
}

func (s *_selector) Enter(rt *gobt.Runtime, node *gobt.Node) {
	s.ahead(rt, node, node.Children)
}

func (s *_selector) Aback(rt *gobt.Runtime, node *gobt.Node) {
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

func (s *_selector) ahead(rt *gobt.Runtime, node *gobt.Node, children []gobt.NodeId) {
	rt.SetStatus(node.Id, "children", children[1:])
	rt.Ahead(children[0], node)
}

func (s *_selector) leave(rt *gobt.Runtime, node *gobt.Node) {
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
