package composite

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestSelector(t *testing.T) {
	var trace []gobt.NodeId

	newTree := func(r1, r2 gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Children: []gobt.NodeId{2, 3},
					Executor: NewSelector(),
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 2)
						return r1
					}),
				},
				3: {
					Id: 3,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 3)
						return r2
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	t1 := newTree(gobt.Failure, gobt.Failure)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Failure, r1.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t2 := newTree(gobt.Failure, gobt.Success)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Success, r2.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t3 := newTree(gobt.Success, gobt.Failure)
	r3 := gobt.New(t3)
	r3.Run()
	assert.Equal(t, gobt.Success, r3.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	trace = []gobt.NodeId{}
	t4 := newTree(gobt.Success, gobt.Success)
	r4 := gobt.New(t4)
	r4.Run()
	assert.Equal(t, gobt.Success, r4.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)
}
