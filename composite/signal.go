package composite

import "github.com/z-hua/gobt"

type _signal struct {
}

func NewSignal() gobt.IExecutor {
	return &_signal{}
}

func (s *_signal) Name() string {
	return "signal"
}

func (s *_signal) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetStatus(node.Id, "flag", struct{}{})
	rt.Ahead(node.Children[0], node)
}

func (s *_signal) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Status(node.Id, "flag") != nil {
		rt.ClrStatus(node.Id)
		if rt.Result().IsSuccess() {
			rt.Ahead(node.Children[1], node)
			return
		}
		rt.SetRunning(node.Id)
		return
	}
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
