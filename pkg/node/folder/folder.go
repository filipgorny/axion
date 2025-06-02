package folder

import "github.com/filipgorny/axion/pkg/node"

type FolderNode struct {
	name     string
	path     *node.Path
	Children []node.Node
}

func NewFolderNode(name string, children ...node.Node) *FolderNode {
	return &FolderNode{
		name:     name,
		Children: append(make([]node.Node, 0), children...),
		path:     node.NewPath(name),
	}
}

func (f *FolderNode) Path() *node.Path {
	if f.path == nil {
		f.path = node.NewPath(f.name)
	}
	return f.path
}

func (f *FolderNode) Name() string {
	return f.name
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

func (f *FolderNode) Find(query *node.NodeQuery) node.Node {
	return f.FindWithPath(query, node.NewPath(query.Name))
}

func (f *FolderNode) FindWithPath(query *node.NodeQuery, path *node.Path) node.Node {
	finalPath := node.NewPath(query.Name)

	for _, child := range f.Children {
		if child.Type() == "folder" {
			folderChild := child.(*FolderNode)

			folderResult := folderChild.FindWithPath(query, finalPath)

			if folderResult != nil {
				return folderResult
			}
		}

		if query.Matches(child) {
			if correctobj, ok := child.(*node.StandardNode); ok {
				correctobj.SetPath(finalPath)

				return correctobj
			}

			return child
		}
	}

	return nil
}

func (f *FolderNode) AddElement(element node.Node) {
	if element == nil {
		return
	}

	if standardElement, ok := element.(*node.StandardNode); ok {
		standardElement.SetPath(node.NewPath(f.Name()))

		element = standardElement
	}

	f.Children = append(f.Children, element)
}

func (f *FolderNode) RemoveElement(element node.Node) {
	for i, child := range f.Children {
		if child == element {
			f.Children = append(f.Children[:i], f.Children[i+1:]...)
			return
		}
	}
}
