package performer

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"testing"
)

func TestNoop(t *testing.T) {
	newTree := func() *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewNoop(),
				},
			},
		}
	}

	t1 := newTree()
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
}
