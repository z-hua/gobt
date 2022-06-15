package composite

import (
	"github.com/z-hua/gobt"
)

type _event struct {
	event gobt.Event
	once  bool
}

func NewEvent(event gobt.Event) gobt.IExecutor {
	return &_event{event: event}
}

func (e *_event) Name() string {
	return "event"
}

func (e *_event) Enter(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Trigger() {
		rt.Ahead(node.Children[1], node)
		return
	}

	rt.Ahead(node.Children[0], node)
	if rt.Result().IsRunning() {
		rt.SetRunning(node.Id)
		rt.SetListen(e.event)
	}
}

func (e *_event) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.Leave(node)
}
