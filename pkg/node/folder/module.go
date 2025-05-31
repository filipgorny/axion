package folder

import "github.com/filipgorny/axion/package/node"

type ModuleNode struct {
	FolderNode
}

func (f *ModuleNode) Type() node.NodeType {
	return "module"
}
