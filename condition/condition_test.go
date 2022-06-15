package condition

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"testing"
)

func TestCondition(t *testing.T) {
	newTree := func(flag bool) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewCondition(func(params ...interface{}) bool { return flag }),
				},
			},
		}
	}

	t1 := newTree(true)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, r1.Result(), gobt.Success)

	t2 := newTree(false)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, r2.Result(), gobt.Failure)
}
