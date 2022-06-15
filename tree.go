package gobt

// Tree 行为树配置数据
type Tree struct {
	Id    TreeId
	Entry NodeId
	Nodes map[NodeId]*Node
}
