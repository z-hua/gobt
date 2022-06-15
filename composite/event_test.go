package composite

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestEvent(t *testing.T) {
	var trace []gobt.NodeId
	var result gobt.Result

	newTree := func() *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewEvent("test"),
					Children: []gobt.NodeId{2, 3},
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 2)
						return result
					}),
				},
				3: {
					Id: 3,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 3)
						return gobt.Success
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	result = gobt.Success
	t1 := newTree()
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	trace = []gobt.NodeId{}
	result = gobt.Failure
	t2 := newTree()
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	trace = []gobt.NodeId{}
	result = gobt.Running
	t3 := newTree()
	r3 := gobt.New(t3)
	r3.Run()
	assert.Equal(t, gobt.Running, r3.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)
	r3.Run()
	assert.Equal(t, gobt.Running, r3.Result())
	assert.Equal(t, []gobt.NodeId{2, 2}, trace)
	result = gobt.Success
	r3.Run()
	assert.Equal(t, gobt.Success, r3.Result())
	assert.Equal(t, []gobt.NodeId{2, 2, 2}, trace)

	trace = []gobt.NodeId{}
	result = gobt.Running
	t4 := newTree()
	r4 := gobt.New(t4)
	r4.Run()
	assert.Equal(t, gobt.Running, r4.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)
	r4.Event("xxx")
	assert.Equal(t, gobt.Running, r4.Result())
	assert.Equal(t, []gobt.NodeId{2, 2}, trace)
	r4.Event("test")
	assert.Equal(t, gobt.Success, r4.Result())
	assert.Equal(t, []gobt.NodeId{2, 2, 3}, trace)
}
