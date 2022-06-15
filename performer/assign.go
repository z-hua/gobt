package performer

import "github.com/z-hua/gobt"

/**
 * 赋值节点
 * key : 变量名
 * val : 变量值
 */
type _assign struct {
	key string
	val int
}

func NewAssign(key string, val int) gobt.IExecutor {
	return &_assign{key: key, val: val}
}

func (a *_assign) Name() string {
	return "assign"
}

func (a *_assign) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetVariant(a.key, a.val)
	rt.SetSuccess()
	rt.Leave(node)
}

func (a *_assign) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
