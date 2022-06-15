package performer

import "github.com/z-hua/gobt"

/**
 * 空节点
 */
type _noop struct {
}

func NewNoop() gobt.IExecutor {
	return &_noop{}
}

func (n *_noop) Name() string {
	return "noop"
}

func (n *_noop) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetSuccess()
	rt.Leave(node)
}

func (n _noop) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
