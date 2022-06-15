package composite

import "github.com/z-hua/gobt"

type _case struct {
}

func NewCase() gobt.IExecutor {
	return &_case{}
}

func (c *_case) Name() string {
	return "case"
}

func (c *_case) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.Ahead(node.Children[0], node)
}

func (c *_case) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.Leave(node)
}
