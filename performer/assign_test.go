package performer

import (
	"github.com/stretchr/testify/assert"
	"github.com/z-hua/gobt"
	"testing"
)

func TestAssign(t *testing.T) {
	newTree := func(key string, val int) *gobt.Tree {
		return &gobt.Tree{
			Entry: 1,
			Nodes: map[gobt.NodeId]*gobt.Node{
				1: {
					Id:       1,
					Executor: NewAssign(key, val),
				},
			},
		}
	}

	t1 := newTree("a", 10)
	r1 := gobt.New(t1)
	r1.Run()
	assert.Equal(t, gobt.Success, r1.Result())
	assert.Equal(t, 10, r1.Variant("a"))
	assert.PanicsWithValue(t, "variant(b) not set", func() {
		r1.Variant("b")
	})
}
