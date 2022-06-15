package decorator

import (
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestLog(t *testing.T) {
	newTree := func() *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewLogger("enter", "aback"),
					Children: []gobt.NodeId{2},
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						return gobt.Success
					}),
				},
			},
		}
	}

	t1 := newTree()
	r1 := gobt.New(t1)
	r1.Run()
}
