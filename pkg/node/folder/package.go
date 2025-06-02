package folder

import "github.com/filipgorny/axion/pkg/node"

type PackageNode struct {
	FolderNode
}

func NewPackageNode(name string, children ...node.Node) *PackageNode {
	return &PackageNode{
		FolderNode: FolderNode{
			name:     name,
			path:     node.NewPath(name),
			Children: append(make([]node.Node, 0), children...),
		},
	}
}

func (f *PackageNode) Type() node.NodeType {
	return "package"
}

func (f *PackageNode) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"it is used to group nodes together in a package",
		"it can contain other folders, components, dtos, etc.",
		"it is often used to represent a package in the application",
	}
}
