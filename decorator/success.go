package decorator

import "github.com/z-hua/gobt"

/**
 * 成功节点
 * 总是返回 success
 */
type _success struct {
}

func NewSuccess() gobt.IExecutor {
	return &_success{}
}

func (s *_success) Name() string {
	return "success"
}

func (s *_success) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.Ahead(node.Children[0], node)
}

func (s *_success) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetSuccess()
	rt.Leave(node)
}
