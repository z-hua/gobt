package decorator

import "github.com/z-hua/gobt"

/**
 * 循环节点
 * 循环执行子节点，并向父节点返回子节点执行结果
 * loop : 循环方式
 * 		-1 : 无限循环
 * 		-2 : 子节点返回 success 时停止
 * 		-3 : 子节点返回 failure 时停止
 * 		 N : 循环 N 次
 * incr : 何时增加计数
 * 		1 : 总是增加
 * 		2 : 子节点返回 success 时增加
 * 		3 : 子节点返回 failure 时增加
 */
type _loop struct {
	loop int
	incr byte
}

func NewLoop(loop int, incr byte) gobt.IExecutor {
	return &_loop{loop: loop, incr: incr}
}

func (l *_loop) Name() string {
	return "loop"
}

func (l *_loop) Enter(rt *gobt.Runtime, node *gobt.Node) {
	if rt.Status(node.Id, "count") == nil {
		rt.SetStatus(node.Id, "count", l.loop)
	}
	rt.Ahead(node.Children[0], node)
}

func (l *_loop) Aback(rt *gobt.Runtime, node *gobt.Node) {
	count := rt.Status(node.Id, "count").(int)
	switch {
	case l.loop == -1:
	case l.loop == -2 && rt.Result().IsSuccess():
		count = 0
	case l.loop == -3 && rt.Result().IsFailure():
		count = 0
	case l.loop > 0:
		switch {
		case l.incr == 1:
			count--
		case l.incr == 2 && rt.Result().IsSuccess():
			count--
		case l.incr == 3 && rt.Result().IsFailure():
			count--
		}
	}
	if count == 0 {
		rt.ClrStatus(node.Id)
		rt.Leave(node)
		return
	}
	rt.SetStatus(node.Id, "count", count)
	rt.SetRunning(node.Id)
}
