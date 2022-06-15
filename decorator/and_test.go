package decorator

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/condition"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestAnd(t *testing.T) {
	var trace []gobt.NodeId

	newTree := func(flag bool, result gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewAnd(),
					Children: []gobt.NodeId{2, 3},
				},
				2: {
					Id: 2,
					Executor: condition.NewCondition(func(params ...interface{}) bool {
						trace = append(trace, 2)
						return flag
					}),
				},
				3: {
					Id: 3,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 3)
						return result
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	t1 := newTree(false, gobt.Success)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Failure, r1.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	trace = []gobt.NodeId{}
	t2 := newTree(false, gobt.Failure)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	trace = []gobt.NodeId{}
	t3 := newTree(true, gobt.Success)
	r3 := gobt.New(t3)
	r3.Run()
	assert.Equal(t, gobt.Success, r3.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t4 := newTree(true, gobt.Failure)
	r4 := gobt.New(t4)
	r4.Run()
	assert.Equal(t, gobt.Failure, r4.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)
}
