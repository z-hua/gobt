package performer

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"testing"
)

func TestAction(t *testing.T) {
	newTree := func(result gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewAction(func(params ...interface{}) gobt.Result { return result }),
				},
			},
		}
	}

	t1 := newTree(gobt.Success)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())

	t2 := newTree(gobt.Failure)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())

	t3 := newTree(gobt.Running)
	r3 := gobt.New(t3)
	r3.Run()
	assert.Equal(t, gobt.Running, r3.Result())
}
