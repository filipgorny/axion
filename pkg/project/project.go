package project

import (
	"github.com/filipgorny/axion/pkg/node/folder"
)

type Project struct {
	Name     string
	RootNode *folder.FolderNode
}

func NewProject(name string) *Project {
	folder := folder.NewFolderNode("root")

	return &Project{
		Name:     name,
		RootNode: folder,
	}
}

func (p *Project) AddModule(module *folder.ModuleNode) {
	p.RootNode.AddElement(module)
}

func (p *Project) GetModule(name string) *folder.ModuleNode {
	for _, child := range p.RootNode.Children {
		if child.Type() == "module" {
			if module, ok := child.(*folder.ModuleNode); ok && module.Name() == name {
				return module
			}
		}
	}

	return nil
}
