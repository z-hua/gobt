package composite

import "github.com/z-hua/gobt"

type _subtree struct {
	subtree *gobt.Tree
	replace bool
}

func NewSubtree(subtree *gobt.Tree, replace bool) gobt.IExecutor {
	return &_subtree{subtree: subtree, replace: replace}
}

func (s *_subtree) Name() string {
	return "subtree"
}

func (s *_subtree) Enter(rt *gobt.Runtime, node *gobt.Node) {
	if s.replace {
		rt.Reset()
		rt.SetTree(s.subtree)
		rt.Ahead(s.subtree.Entry, nil)
		return
	}
	rt.SetTree(s.subtree)
	rt.Ahead(s.subtree.Entry, node)
}

func (s *_subtree) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.Leave(node)
}
