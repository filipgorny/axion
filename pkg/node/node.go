package node

type NodeType string

type Node interface {
	Type() NodeType
}
