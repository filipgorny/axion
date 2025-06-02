package folder

import "github.com/filipgorny/axion/pkg/node"

type ModuleNode struct {
	FolderNode
}

func NewModuleNode(name string, children ...node.Node) *ModuleNode {
	return &ModuleNode{
		FolderNode: FolderNode{
			name:     name,
			path:     node.NewPath(name),
			Children: append(make([]node.Node, 0), children...),
		},
	}
}

func (f *ModuleNode) Type() node.NodeType {
	return "module"
}

func (f *ModuleNode) Characteristics() node.NodeCharacteristics {
	return node.Characteristics(
		"it is used to group nodes together in a module",
		"it can contain other folders, components, dtos, etc.",
		"it is often used to represent a feature or a domain in the application",
	)
}
