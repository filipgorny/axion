package component

import "github.com/filipgorny/axion/pkg/node"

type ComponentNode struct {
	Name string
}

func NewComponentNode(name string) *ComponentNode {
	return &ComponentNode{
		Name: name,
	}
}

func (c *ComponentNode) Type() node.NodeType {
	return "component"
}

func (c *ComponentNode) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"it is often used to represent a html portion of the view",
		"it is sometimes reusable, othertimes it is a main component like a page",
	}
}
