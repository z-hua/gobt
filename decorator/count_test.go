package decorator

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestCount(t *testing.T) {
	var trace []gobt.NodeId
	var result gobt.Result

	newTree := func(times int, incr byte) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewCount(times, incr),
					Children: []gobt.NodeId{2},
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 2)
						return result
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	result = gobt.Success
	t1 := newTree(2, 1)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
	assert.Equal(t, []gobt.NodeId{2, 2}, trace)
	r1.Run()
	assert.Equal(t, gobt.Failure, r1.Result())
	assert.Equal(t, []gobt.NodeId{2, 2}, trace)

	trace = []gobt.NodeId{}
	result = gobt.Failure
	t2 := newTree(1, 2)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)
	result = gobt.Success
	r2.Run()
	assert.Equal(t, gobt.Success, r2.Result())
	assert.Equal(t, []gobt.NodeId{2, 2}, trace)
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())
	assert.Equal(t, []gobt.NodeId{2, 2}, trace)
}
