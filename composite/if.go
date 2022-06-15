package composite

import "github.com/z-hua/gobt"

/**
 * if节点
 * 相当于if-else语句，有2/3个子节点
 * 第 1 个子节点为条件判断节点
 * 第 2 个子节点在条件判断成功时执行
 * 第 3 个子节点在条件判断失败时执行(可选)
 */
type _if struct {
}

func NewIf() gobt.IExecutor {
	return &_if{}
}

func (i *_if) Name() string {
	return "if"
}

func (i *_if) Enter(rt *gobt.Runtime, node *gobt.Node) {
	rt.SetStatus(node.Id, "flag", struct{}{})
	rt.Ahead(node.Children[0], node)
}

func (i *_if) Aback(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Status(node.Id, "flag") != nil {
		rt.ClrStatus(node.Id)
		if rt.Result().IsSuccess() {
			rt.Ahead(node.Children[1], node)
			return
		}
		if len(node.Children) == 3 {
			rt.Ahead(node.Children[2], node)
			return
		}
	}
	rt.ClrStatus(node.Id)
	rt.Leave(node)
}
