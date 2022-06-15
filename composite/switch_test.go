package composite

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestSwitch(t *testing.T) {
	var trace []gobt.NodeId

	newTree1 := func(result1, result2 gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewSwitch(),
					Children: []gobt.NodeId{2, 3},
				},
				2: {
					Id:       2,
					Executor: NewCase(),
					Children: []gobt.NodeId{4},
				},
				3: {
					Id:       3,
					Executor: NewCase(),
					Children: []gobt.NodeId{5},
				},
				4: {
					Id: 4,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 4)
						return result1
					}),
				},
				5: {
					Id: 5,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 5)
						return result2
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	t11 := newTree1(gobt.Success, gobt.Success)
	r11 := gobt.New(t11)
	r11.Run()
	assert.Equal(t, gobt.Success, r11.Result())
	assert.Equal(t, []gobt.NodeId{4}, trace)

	trace = []gobt.NodeId{}
	t12 := newTree1(gobt.Failure, gobt.Success)
	r12 := gobt.New(t12)
	r12.Run()
	assert.Equal(t, gobt.Success, r12.Result())
	assert.Equal(t, []gobt.NodeId{4, 5}, trace)

	trace = []gobt.NodeId{}
	t13 := newTree1(gobt.Success, gobt.Failure)
	r13 := gobt.New(t13)
	r13.Run()
	assert.Equal(t, gobt.Success, r13.Result())
	assert.Equal(t, []gobt.NodeId{4}, trace)

	trace = []gobt.NodeId{}
	t14 := newTree1(gobt.Failure, gobt.Failure)
	r14 := gobt.New(t14)
	r14.Run()
	assert.Equal(t, gobt.Failure, r14.Result())
	assert.Equal(t, []gobt.NodeId{4, 5}, trace)

	newTree2 := func(result1, result2, result3 gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewSwitch(),
					Children: []gobt.NodeId{2, 3, 4},
				},
				2: {
					Id:       2,
					Executor: NewCase(),
					Children: []gobt.NodeId{5},
				},
				3: {
					Id:       3,
					Executor: NewCase(),
					Children: []gobt.NodeId{6},
				},
				4: {
					Id: 4,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 4)
						return result3
					}),
				},
				5: {
					Id: 5,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 5)
						return result1
					}),
				},
				6: {
					Id: 6,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 6)
						return result2
					}),
				},
			},
		}
	}

	trace = []gobt.NodeId{}
	t21 := newTree2(gobt.Success, gobt.Success, gobt.Success)
	r21 := gobt.New(t21)
	r21.Run()
	assert.Equal(t, gobt.Success, r21.Result())
	assert.Equal(t, []gobt.NodeId{5}, trace)

	trace = []gobt.NodeId{}
	t22 := newTree2(gobt.Failure, gobt.Success, gobt.Success)
	r22 := gobt.New(t22)
	r22.Run()
	assert.Equal(t, gobt.Success, r22.Result())
	assert.Equal(t, []gobt.NodeId{5, 6}, trace)

	trace = []gobt.NodeId{}
	t23 := newTree2(gobt.Success, gobt.Failure, gobt.Success)
	r23 := gobt.New(t23)
	r23.Run()
	assert.Equal(t, gobt.Success, r23.Result())
	assert.Equal(t, []gobt.NodeId{5}, trace)

	trace = []gobt.NodeId{}
	t24 := newTree2(gobt.Failure, gobt.Failure, gobt.Success)
	r24 := gobt.New(t24)
	r24.Run()
	assert.Equal(t, gobt.Success, r24.Result())
	assert.Equal(t, []gobt.NodeId{5, 6, 4}, trace)

	trace = []gobt.NodeId{}
	t25 := newTree2(gobt.Failure, gobt.Failure, gobt.Failure)
	r25 := gobt.New(t25)
	r25.Run()
	assert.Equal(t, gobt.Failure, r25.Result())
	assert.Equal(t, []gobt.NodeId{5, 6, 4}, trace)
}
