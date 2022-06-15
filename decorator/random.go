package decorator

import (
	"github.com/z-hua/gobt"
	"math/rand"
)

/**
 * 随机节点
 */
type _random struct {
}

func NewRandom() gobt.IExecutor {
	return &_random{}
}

func (r *_random) Name() string {
	return "random"
}

func (r *_random) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.Ahead(node.Children[rand.Intn(len(node.Children))], node)
}

func (r *_random) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.Leave(node)
}
