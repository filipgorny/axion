package folder

import "github.com/filipgorny/axion/package/node"

type PackageNode struct {
	FolderNode
}

func (f *PackageNode) Type() node.NodeType {
	return "package"
}
