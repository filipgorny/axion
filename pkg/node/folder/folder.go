package folder

import "github.com/filipgorny/axion/package/node"

type FolderNode struct {
	Name     string
	Children []*node.Node
}

func (f *FolderNode) Type() node.NodeType {
	return "folder"
}
