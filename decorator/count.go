package decorator

import "github.com/z-hua/gobt"

/**
 * 计数节点
 * times : 节点运行次数
 * incr : 何时增加计数
 * 		1 : 总是增加
 * 		2 : 子节点返回 success 时增加
 * 		3 : 子节点返回 failure 时增加
 */
type _count struct {
	times int
	incr  byte
}

func NewCount(times int, incr byte) gobt.IExecutor {
	return &_count{times: times, incr: incr}
}

func (c *_count) Name() string {
	return "count"
}

func (c *_count) Enter(rt *gobt.Runtime, node *gobt.Node) {
	count, ok := rt.Status(node.Id, "count").(int)
	if !ok {
		count = 0
		rt.SetStatus(node.Id, "count", count)
	}
	if count >= c.times {
		rt.SetFailure()
		rt.Leave(node)
		return
	}
	rt.Ahead(node.Children[0], node)

}

func (c *_count) Aback(rt *gobt.Runtime, node *gobt.Node) {
	count := rt.Status(node.Id, "count").(int)
	switch {
	case c.incr == 1:
		count++
	case c.incr == 2 && rt.Result().IsSuccess():
		count++
	case c.incr == 3 && rt.Result().IsFailure():
		count++
	}
	rt.SetStatus(node.Id, "count", count)
	rt.Leave(node)
}
