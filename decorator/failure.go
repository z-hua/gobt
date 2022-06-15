package decorator

import "github.com/z-hua/gobt"

/**
 * 失败节点
 * 总是返回 failure
 */
type _failure struct {
}

func NewFailure() gobt.IExecutor {
	return &_failure{}
}

func (f *_failure) Name() string {
	return "failure"
}

func (f *_failure) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.Ahead(node.Children[0], node)
}

func (f *_failure) Aback(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetFailure()
	rt.Leave(node)
}
