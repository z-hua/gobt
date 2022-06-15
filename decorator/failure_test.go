package decorator

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/performer"
	"testing"
)

func TestFailure(t *testing.T) {
	newTree := func(result gobt.Result) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewFailure(),
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

	t1 := newTree(gobt.Success)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Failure, r1.Result())

	t2 := newTree(gobt.Failure)
	r2 := gobt.New(t2)
	r2.Run()
	assert.Equal(t, gobt.Failure, r2.Result())
}
