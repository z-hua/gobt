package composite

import "github.com/z-hua/gobt"

/**
 * 并行节点
 * 执行所有子节点, 根据 types 属性向父节点返回结果
 * types
 * 	 1 : 有一个子节点返回 success, 就向父节点返回 success, 否则返回 failure
 *   2 : 所有子节点都返回 success, 才向父节点返回 success, 否则返回 failure
 *   3 : 有一个子节点返回 failure, 就向父节点返回 failure, 否则返回 success
 *   4 : 所有子节点才返回 failure, 才向父节点返回 failure, 否则返回 success
 */
type _parallel struct {
	types byte
}

func NewParallel(types byte) gobt.IExecutor {
	return &_parallel{types: types}
}

func (p *_parallel) Name() string {
	return "parallel"
}

func (p *_parallel) Enter(rt *gobt.Runtime, node *gobt.Node) {
	var result gobt.Result
	switch p.types {
	case 1, 4:
		result = gobt.Failure
	case 2, 3:
		result = gobt.Success
	}
	p.ahead(rt, node, node.Children, result)
}

func (p *_parallel) Aback(rt *gobt.Runtime, node *gobt.Node) {
	children := rt.Status(node.Id, "children").([]gobt.NodeId)
	result := rt.Status(node.Id, "result").(gobt.Result)
	switch {
	case p.types == 1 && rt.Result().IsSuccess():
		result = gobt.Success
	case p.types == 2 && rt.Result().IsFailure():
		result = gobt.Failure
	case p.types == 3 && rt.Result().IsFailure():
		result = gobt.Failure
	case p.types == 4 && rt.Result().IsSuccess():
		result = gobt.Success
	}
	if len(children) == 0 {
		p.leave(rt, node, result)
		return
	}
	p.ahead(rt, node, children, result)
}

func (p *_parallel) ahead(rt *gobt.Runtime, node *gobt.Node, children []gobt.NodeId, result gobt.Result) {
	rt.SetStatus(node.Id, "result", result)
	rt.SetStatus(node.Id, "children", children[1:])
	rt.Ahead(children[0], node)
}

func (p *_parallel) leave(rt *gobt.Runtime, node *gobt.Node, result gobt.Result) {
	rt.ClrStatus(node.Id)
	rt.SetResult(result)
	rt.Leave(node)
}
