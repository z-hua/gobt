package decorator

import (
	"github.com/z-hua/gobt"
	"math/rand"
)

/**
 * 权重节点
 */
type _weight struct {
	weights []int
	total   int
}

func NewWeight(weights []int) gobt.IExecutor {
	w := &_weight{weights: weights}
	for _, v := range weights {
		w.total += v
	}
	return w
}

func (w *_weight) Name() string {
	return "weight"
}

func (w *_weight) Enter(rt *gobt.Runtime, node *gobt.Node) {
	var idx, acc int
	hit := rand.Intn(w.total) + 1
	for i, e := range w.weights {
		acc += e
		if hit <= acc {
			idx = i
			break
		}
	}
	rt.Ahead(node.Children[idx], node)
}

func (w *_weight) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.Leave(node)
}
