package performer

import "github.com/z-hua/gobt"

type _break struct {
	whole bool
}

func NewBreak(whole bool) gobt.IExecutor {
	return &_break{whole: whole}
}

func (b *_break) Name() string {
	return "break"
}

func (b *_break) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.Abort(b.whole)
	rt.Leave(node)
}

func (b *_break) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
