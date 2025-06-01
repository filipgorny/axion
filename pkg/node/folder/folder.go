package folder

import "github.com/filipgorny/axion/pkg/node"

type FolderNode struct {
	Name     string
	Children []*node.Node
}

func NewFolderNode(name string, children ...*node.Node) *FolderNode {
	return &FolderNode{
		Name:     name,
		Children: append(make([]*node.Node, 0), children...),
	}
}

func (f *FolderNode) Type() node.NodeType {
	return "folder"
}

func (f *FolderNode) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"it is used to group nodes together",
		"it can contain other folders, components, dtos, etc.",
	}
}
