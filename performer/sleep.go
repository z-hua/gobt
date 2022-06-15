package performer

import "github.com/z-hua/gobt"

type _sleep struct {
	frame int
	yield bool
}

func NewSleep(frame int, yield bool) gobt.IExecutor {
	return &_sleep{frame: frame, yield: yield}
}

func (s *_sleep) Name() string {
	return "sleep"
}

func (s *_sleep) Enter(rt *gobt.Runtime, node *gobt.Node) {
	count, ok := rt.Status(node.Id, "count").(int)
	if !ok {
		count = s.frame
	} else {
		count--
	}
	if count <= 0 {
		rt.SetSuccess()
		rt.ClrStatus(node.Id)
		rt.Leave(node)
		return
	}
	rt.SetRunning(node.Id)
	rt.SetStatus(node.Id, "count", count)
	rt.Leave(node)
}

func (s *_sleep) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
