package composite

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/condition"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestIf(t *testing.T) {
	var trace []gobt.NodeId

	newTree1 := func(flag bool, result gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewIf(),
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
	t11 := newTree1(true, gobt.Success)
	r11 := gobt.New(t11)
	r11.Run()
	assert.Equal(t, gobt.Success, r11.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t12 := newTree1(true, gobt.Failure)
	r12 := gobt.New(t12)
	r12.Run()
	assert.Equal(t, gobt.Failure, r12.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t13 := newTree1(false, gobt.Success)
	r13 := gobt.New(t13)
	r13.Run()
	assert.Equal(t, gobt.Failure, r13.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	trace = []gobt.NodeId{}
	t14 := newTree1(false, gobt.Failure)
	r14 := gobt.New(t14)
	r14.Run()
	assert.Equal(t, gobt.Failure, r14.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)

	newTree2 := func(flag bool, result1, result2 gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewIf(),
					Children: []gobt.NodeId{2, 3, 4},
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
						return result1
					}),
				},
				4: {
					Id: 4,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 4)
						return result2
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	t21 := newTree2(true, gobt.Success, gobt.Success)
	r21 := gobt.New(t21)
	r21.Run()
	assert.Equal(t, gobt.Success, r21.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t22 := newTree2(true, gobt.Success, gobt.Failure)
	r22 := gobt.New(t22)
	r22.Run()
	assert.Equal(t, gobt.Success, r22.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t23 := newTree2(true, gobt.Failure, gobt.Success)
	r23 := gobt.New(t23)
	r23.Run()
	assert.Equal(t, gobt.Failure, r23.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t24 := newTree2(true, gobt.Failure, gobt.Failure)
	r24 := gobt.New(t24)
	r24.Run()
	assert.Equal(t, gobt.Failure, r24.Result())
	assert.Equal(t, []gobt.NodeId{2, 3}, trace)

	trace = []gobt.NodeId{}
	t25 := newTree2(false, gobt.Success, gobt.Success)
	r25 := gobt.New(t25)
	r25.Run()
	assert.Equal(t, gobt.Success, r25.Result())
	assert.Equal(t, []gobt.NodeId{2, 4}, trace)

	trace = []gobt.NodeId{}
	t26 := newTree2(false, gobt.Success, gobt.Failure)
	r26 := gobt.New(t26)
	r26.Run()
	assert.Equal(t, gobt.Failure, r26.Result())
	assert.Equal(t, []gobt.NodeId{2, 4}, trace)

	trace = []gobt.NodeId{}
	t27 := newTree2(false, gobt.Failure, gobt.Success)
	r27 := gobt.New(t27)
	r27.Run()
	assert.Equal(t, gobt.Success, r27.Result())
	assert.Equal(t, []gobt.NodeId{2, 4}, trace)

	trace = []gobt.NodeId{}
	t28 := newTree2(false, gobt.Failure, gobt.Failure)
	r28 := gobt.New(t28)
	r28.Run()
	assert.Equal(t, gobt.Failure, r28.Result())
	assert.Equal(t, []gobt.NodeId{2, 4}, trace)
}
