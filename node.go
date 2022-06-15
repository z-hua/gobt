package gobt

// IExecutor 节点执行器
type IExecutor interface {
	// Name 执行器名称
	Name() string

	// Enter 执行节点时回调
	Enter(rt *Runtime, node *Node)

	// Aback 返回节点时回调
	Aback(rt *Runtime, node *Node)
}

// Node 行为树节点
type Node struct {
	Id       NodeId      // 节点 id
	Children []NodeId    // 子节点列表
	PreCheck func() bool // 前置检查
	Postpone func()      // 后置效果
	Executor IExecutor   // 节点执行器
}
