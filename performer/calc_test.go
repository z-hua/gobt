package performer

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/composite"
	"testing"
)

func TestCalc(t *testing.T) {
	newTree := func(op string, val int) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: composite.NewSequence(),
					Children: []gobt.NodeId{2, 3},
				},
				2: {
					Id:       2,
					Executor: NewAssign("a", 10),
				},
				3: {
					Id:       3,
					Executor: NewCalc("a", op, val),
				},
			},
		}
	}

	t1 := newTree("+", 5)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
	assert.Equal(t, 15, r1.Variant("a"))

	t2 := newTree("-", 5)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Success, r2.Result())
	assert.Equal(t, 5, r2.Variant("a"))

	t3 := newTree("*", 5)
	r3 := gobt.New(t3)
	r3.Run()
	assert.Equal(t, gobt.Success, r3.Result())
	assert.Equal(t, 50, r3.Variant("a"))

	t4 := newTree("/", 5)
	r4 := gobt.New(t4)
	r4.Run()
	assert.Equal(t, gobt.Success, r4.Result())
	assert.Equal(t, 2, r4.Variant("a"))

	t5 := newTree("%", 5)
	r5 := gobt.New(t5)
	r5.Run()
	assert.Equal(t, gobt.Success, r5.Result())
	assert.Equal(t, 0, r5.Variant("a"))

	t6 := newTree("%", 3)
	r6 := gobt.New(t6)
	r6.Run()
	assert.Equal(t, gobt.Success, r6.Result())
	assert.Equal(t, 1, r6.Variant("a"))
}
