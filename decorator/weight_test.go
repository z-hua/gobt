package decorator

import (
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestWeight(t *testing.T) {
	newTree := func() *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewWeight([]int{10, 10}),
					Children: []gobt.NodeId{2, 3},
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						t.Logf("hit node 2")
						return gobt.Success
					}),
				},
				3: {
					Id: 3,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						t.Logf("hit node 3")
						return gobt.Success
					}),
				},
			},
		}
	}

	t1 := newTree()
	r1 := gobt.New(t1)
	for i := 0; i < 10; i++ {
		r1.Run()
	}
}
