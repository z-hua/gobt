package condition

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/composite"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestCompare(t *testing.T) {
	newTree := func(key, cmp string, val1, val2 int) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Children: []gobt.NodeId{2, 3},
					Executor: composite.NewSequence(),
				},
				2: {
					Id:       2,
					Executor: performer.NewAssign(key, val1),
				},
				3: {
					Id:       3,
					Executor: NewCompare(key, cmp, val2),
				},
			},
		}
	}

	t1 := newTree("a", ">", 3, 1)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, r1.Result(), gobt.Success)

	t2 := newTree("a", ">", 3, 5)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, r2.Result(), gobt.Failure)

	t3 := newTree("a", "<", 3, 1)
	r3 := gobt.New(t3)
	r3.Run()
	assert.Equal(t, r3.Result(), gobt.Failure)

	t4 := newTree("a", "<", 3, 5)
	r4 := gobt.New(t4)
	r4.Run()
	assert.Equal(t, r4.Result(), gobt.Success)

	t5 := newTree("a", ">=", 3, 1)
	r5 := gobt.New(t5)
	r5.Run()
	assert.Equal(t, r5.Result(), gobt.Success)

	t6 := newTree("a", ">=", 3, 3)
	r6 := gobt.New(t6)
	r6.Run()
	assert.Equal(t, r6.Result(), gobt.Success)

	t7 := newTree("a", ">=", 3, 5)
	r7 := gobt.New(t7)
	r7.Run()
	assert.Equal(t, r7.Result(), gobt.Failure)

	t8 := newTree("a", "<=", 3, 1)
	r8 := gobt.New(t8)
	r8.Run()
	assert.Equal(t, r8.Result(), gobt.Failure)

	t9 := newTree("a", "<=", 3, 3)
	r9 := gobt.New(t9)
	r9.Run()
	assert.Equal(t, r9.Result(), gobt.Success)

	t10 := newTree("a", "<=", 3, 5)
	r10 := gobt.New(t10)
	r10.Run()
	assert.Equal(t, r10.Result(), gobt.Success)

	t11 := newTree("a", "==", 3, 3)
	r11 := gobt.New(t11)
	r11.Run()
	assert.Equal(t, r11.Result(), gobt.Success)

	t12 := newTree("a", "==", 3, 5)
	r12 := gobt.New(t12)
	r12.Run()
	assert.Equal(t, r12.Result(), gobt.Failure)

	t13 := newTree("a", "!=", 3, 3)
	r13 := gobt.New(t13)
	r13.Run()
	assert.Equal(t, r13.Result(), gobt.Failure)

	t14 := newTree("a", "!=", 3, 5)
	r14 := gobt.New(t14)
	r14.Run()
	assert.Equal(t, r14.Result(), gobt.Success)
}
