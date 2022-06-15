package gobt

import "fmt"

type (
	TreeId int    // 行为树 id
	RunRef int    // 运行时引用
	NodeId int    // 节点 id
	Event  string // 事件
)

// Result 运行结果
type Result int

func (r Result) IsSuccess() bool {
	return r == Success
}

func (r Result) IsFailure() bool {
	return r == Failure
}

func (r Result) IsRunning() bool {
	return r == Running
}

const (
	Success Result = iota
	Failure
	Running
)

type from struct {
	tree *Tree
	node *Node
}

// Runtime 行为树运行时数据
type Runtime struct {
	tree    *Tree
	params  []interface{}                     // 运行参数
	result  Result                            // 运行结果
	running NodeId                            // 运行中的节点
	tracing []*from                           // 节点执行路径
	status  map[NodeId]map[string]interface{} // 节点状态
	variant map[string]int                    // 变量定义
	accept  bool                              // 是否接受事件
	listen  Event                             // 正在监听的事件
	trigger bool                              // 事件是否触发
}

func New(tree *Tree) *Runtime {
	return &Runtime{
		tree:    tree,
		tracing: make([]*from, 0, 3),
		status:  map[NodeId]map[string]interface{}{},
		variant: map[string]int{},
	}
}

func (rt *Runtime) Run(params ...interface{}) {
	rt.params = params

	if rt.running == 0 {
		rt.Ahead(rt.tree.Entry, nil)
		return
	}

	if fr := rt.popFrom(); fr != nil {
		rt.Ahead(rt.running, fr.node)
		return
	}

	rt.Ahead(rt.running, nil)
}

func (rt *Runtime) Event(event Event, params ...interface{}) {
	if rt.listen == event {
		rt.trigger = true
		rt.Run(params...)
		rt.trigger = false
		return
	}
	rt.Run(params...)
}

func (rt *Runtime) Reset() {
	rt.tracing = make([]*from, 0, 3)
	rt.status = map[NodeId]map[string]interface{}{}
	rt.variant = map[string]int{}
}

func (rt *Runtime) Abort(whole bool) {
	tracing := make([]*from, 0, 3)
	if !whole {
		for _, fr := range rt.tracing {
			if fr.tree.Id != rt.tree.Id {
				tracing = append(tracing, fr)
			}
		}
	}
	rt.tracing = tracing
}

func (rt *Runtime) Ahead(next NodeId, from *Node) {
	node, ok := rt.tree.Nodes[next]
	if !ok {
		panic("node not exist")
	}
	if node.PreCheck != nil && !node.PreCheck() {
		rt.SetFailure()
		return
	}
	if from != nil {
		rt.pushFrom(from)
	}
	node.Executor.Enter(rt, node)
}

func (rt *Runtime) Leave(node *Node) {
	if node.Postpone != nil {
		node.Postpone()
	}
	if fr := rt.popFrom(); fr != nil {
		rt.tree = fr.tree
		fr.node.Executor.Aback(rt, fr.node)
	}
}

func (rt *Runtime) Params() []interface{} {
	return rt.params
}

func (rt *Runtime) Result() Result {
	return rt.result
}

func (rt *Runtime) SetResult(result Result) {
	rt.result = result
}

func (rt *Runtime) SetSuccess() {
	rt.result = Success
}

func (rt *Runtime) SetFailure() {
	rt.result = Failure
}

func (rt *Runtime) SetRunning(running NodeId) {
	rt.result = Running
	rt.running = running
}

func (rt *Runtime) Status(id NodeId, key string) interface{} {
	if m, ok := rt.status[id]; ok {
		return m[key]
	}
	return nil
}

func (rt *Runtime) SetStatus(id NodeId, key string, val interface{}) {
	m, ok := rt.status[id]
	if !ok {
		m = map[string]interface{}{}
	}
	m[key] = val
	rt.status[id] = m
}

func (rt *Runtime) ClrStatus(id NodeId) {
	rt.status[id] = map[string]interface{}{}
}

func (rt *Runtime) Variant(key string) int {
	val, ok := rt.variant[key]
	if !ok {
		panic(fmt.Sprintf("variant(%s) not set", key))
	}
	return val
}

func (rt *Runtime) SetVariant(key string, val int) {
	rt.variant[key] = val
}

func (rt *Runtime) SetTree(tree *Tree) {
	rt.tree = tree
}

func (rt *Runtime) Listen() Event {
	return rt.listen
}

func (rt *Runtime) SetListen(event Event) {
	rt.listen = event
}

func (rt *Runtime) Trigger() bool {
	return rt.trigger
}

func (rt *Runtime) pushFrom(node *Node) {
	rt.tracing = append(rt.tracing, &from{tree: rt.tree, node: node})
}

func (rt *Runtime) popFrom() *from {
	length := len(rt.tracing)
	if length == 0 {
		return nil
	}
	result := rt.tracing[length-1]
	rt.tracing = rt.tracing[:length-1]
	return result
}

func (rt *Runtime) clrFrom() {
	rt.tracing = make([]*from, 0, 5)
}
