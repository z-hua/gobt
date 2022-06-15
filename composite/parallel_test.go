package composite

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestParallel(t *testing.T) {
	newTree := func(types byte, r1, r2 gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Children: []gobt.NodeId{2, 3},
					Executor: NewParallel(types),
				},
				2: {
					Id: 2,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						return r1
					}),
				},
				3: {
					Id: 3,
					Executor: performer.NewAction(func(params ...interface{}) gobt.Result {
						return r2
					}),
				},
			},
		}
	}

	t11 := newTree(1, gobt.Failure, gobt.Failure)
	r11 := gobt.New(t11)
	r11.Run()
	assert.Equal(t, gobt.Failure, r11.Result())

	t12 := newTree(1, gobt.Success, gobt.Failure)
	r12 := gobt.New(t12)
	r12.Run()
	assert.Equal(t, gobt.Success, r12.Result())

	t13 := newTree(1, gobt.Failure, gobt.Success)
	r13 := gobt.New(t13)
	r13.Run()
	assert.Equal(t, gobt.Success, r13.Result())

	t14 := newTree(1, gobt.Success, gobt.Success)
	r14 := gobt.New(t14)
	r14.Run()
	assert.Equal(t, gobt.Success, r14.Result())

	t21 := newTree(2, gobt.Failure, gobt.Failure)
	r21 := gobt.New(t21)
	r21.Run()
	assert.Equal(t, gobt.Failure, r21.Result())

	t22 := newTree(2, gobt.Success, gobt.Failure)
	r22 := gobt.New(t22)
	r22.Run()
	assert.Equal(t, gobt.Failure, r22.Result())

	t23 := newTree(2, gobt.Failure, gobt.Success)
	r23 := gobt.New(t23)
	r23.Run()
	assert.Equal(t, gobt.Failure, r23.Result())

	t24 := newTree(2, gobt.Success, gobt.Success)
	r24 := gobt.New(t24)
	r24.Run()
	assert.Equal(t, gobt.Success, r24.Result())

	t31 := newTree(3, gobt.Failure, gobt.Failure)
	r31 := gobt.New(t31)
	r31.Run()
	assert.Equal(t, gobt.Failure, r31.Result())

	t32 := newTree(3, gobt.Success, gobt.Failure)
	r32 := gobt.New(t32)
	r32.Run()
	assert.Equal(t, gobt.Failure, r32.Result())

	t33 := newTree(3, gobt.Failure, gobt.Success)
	r33 := gobt.New(t33)
	r33.Run()
	assert.Equal(t, gobt.Failure, r33.Result())

	t34 := newTree(3, gobt.Success, gobt.Success)
	r34 := gobt.New(t34)
	r34.Run()
	assert.Equal(t, gobt.Success, r34.Result())

	t41 := newTree(4, gobt.Failure, gobt.Failure)
	r41 := gobt.New(t41)
	r41.Run()
	assert.Equal(t, gobt.Failure, r41.Result())

	t42 := newTree(4, gobt.Success, gobt.Failure)
	r42 := gobt.New(t42)
	r42.Run()
	assert.Equal(t, gobt.Success, r42.Result())

	t43 := newTree(4, gobt.Failure, gobt.Success)
	r43 := gobt.New(t43)
	r43.Run()
	assert.Equal(t, gobt.Success, r43.Result())

	t44 := newTree(4, gobt.Success, gobt.Success)
	r44 := gobt.New(t44)
	r44.Run()
	assert.Equal(t, gobt.Success, r44.Result())
}
