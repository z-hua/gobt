package composite

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestSubtree(t *testing.T) {
	var trace []gobt.NodeId

	subTree := func(result gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 10,
			Nodes: map[gobt.NodeId]*gobt.Node{
				10: {
					Id: 10,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 10)
						return result
					}),
				},
			},
		}
	}
	newTree := func(replace bool, result1, result2 gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewSequence(),
					Children: []gobt.NodeId{2, 3},
				},
				2: {
					Id:       2,
					Executor: NewSubtree(subTree(result1), replace),
				},
				3: {
					Id: 3,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 3)
						return result2
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	t11 := newTree(false, gobt.Success, gobt.Success)
	r11 := gobt.New(t11)
	r11.Run()
	assert.Equal(t, gobt.Success, r11.Result())
	assert.Equal(t, []gobt.NodeId{10, 3}, trace)

	trace = []gobt.NodeId{}
	t12 := newTree(false, gobt.Success, gobt.Failure)
	r12 := gobt.New(t12)
	r12.Run()
	assert.Equal(t, gobt.Failure, r12.Result())
	assert.Equal(t, []gobt.NodeId{10, 3}, trace)

	trace = []gobt.NodeId{}
	t13 := newTree(false, gobt.Failure, gobt.Success)
	r13 := gobt.New(t13)
	r13.Run()
	assert.Equal(t, gobt.Failure, r13.Result())
	assert.Equal(t, []gobt.NodeId{10}, trace)

	trace = []gobt.NodeId{}
	t14 := newTree(false, gobt.Failure, gobt.Failure)
	r14 := gobt.New(t14)
	r14.Run()
	assert.Equal(t, gobt.Failure, r14.Result())
	assert.Equal(t, []gobt.NodeId{10}, trace)

	trace = []gobt.NodeId{}
	t21 := newTree(true, gobt.Success, gobt.Success)
	r21 := gobt.New(t21)
	r21.Run()
	assert.Equal(t, gobt.Success, r21.Result())
	assert.Equal(t, []gobt.NodeId{10}, trace)

	trace = []gobt.NodeId{}
	t22 := newTree(true, gobt.Failure, gobt.Failure)
	r22 := gobt.New(t22)
	r22.Run()
	assert.Equal(t, gobt.Failure, r22.Result())
	assert.Equal(t, []gobt.NodeId{10}, trace)
}
