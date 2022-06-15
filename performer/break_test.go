package performer

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"github.com/z-hua/gobt/composite"
	"testing"
)

func TestBreak(t *testing.T) {
	var trace []gobt.NodeId

	newTree := func() *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: composite.NewSequence(),
					Children: []gobt.NodeId{2, 3, 4},
				},
				2: {
					Id: 2,
					Executor: NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 2)
						return gobt.Success
					}),
				},
				3: {
					Id:       3,
					Executor: NewBreak(true),
				},
				4: {
					Id: 4,
					Executor: NewAction(func(params ...interface{}) gobt.Result {
						trace = append(trace, 4)
						return gobt.Success
					}),
				},
			},
		}
	}

	t1 := newTree()
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
	assert.Equal(t, []gobt.NodeId{2}, trace)
}
