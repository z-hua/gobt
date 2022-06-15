package decorator

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestLoop(t *testing.T) {
	var result gobt.Result
	newTree := func(loop int, incr byte) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewLoop(loop, incr),
					Children: []gobt.NodeId{2},
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						return result
					}),
				},
			},
		}
	}

	t1 := newTree(-2, 0)
	r1 := gobt.New(t1)
	result = gobt.Failure
	r1.Run()
	assert.Equal(t, gobt.Running, r1.Result())
	r1.Run()
	assert.Equal(t, gobt.Running, r1.Result())
	result = gobt.Success
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())

	t2 := newTree(-3, 0)
	r2 := gobt.New(t2)
	result = gobt.Success
	r2.Run()
	assert.Equal(t, gobt.Running, r2.Result())
	r2.Run()
	assert.Equal(t, gobt.Running, r2.Result())
	result = gobt.Failure
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())

	t3 := newTree(3, 1)
	r3 := gobt.New(t3)
	result = gobt.Success
	r3.Run()
	assert.Equal(t, gobt.Running, r3.Result())
	r3.Run()
	assert.Equal(t, gobt.Running, r3.Result())
	r3.Run()
	assert.Equal(t, gobt.Success, r3.Result())

	t4 := newTree(3, 2)
	r4 := gobt.New(t4)
	result = gobt.Failure
	r4.Run()
	assert.Equal(t, gobt.Running, r4.Result())
	result = gobt.Success
	r4.Run()
	assert.Equal(t, gobt.Running, r4.Result())
	r4.Run()
	assert.Equal(t, gobt.Running, r4.Result())
	r4.Run()
	assert.Equal(t, gobt.Success, r4.Result())

	t5 := newTree(3, 3)
	r5 := gobt.New(t5)
	result = gobt.Success
	r5.Run()
	assert.Equal(t, gobt.Running, r5.Result())
	result = gobt.Failure
	r5.Run()
	assert.Equal(t, gobt.Running, r5.Result())
	r5.Run()
	assert.Equal(t, gobt.Running, r5.Result())
	r5.Run()
	assert.Equal(t, gobt.Failure, r5.Result())
}
