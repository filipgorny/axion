package action

import (
	"github.com/filipgorny/axion/pkg/node"
	"github.com/filipgorny/axion/pkg/node/folder"
	"github.com/filipgorny/axion/pkg/project"
)

type ProjectAddModule struct {
	parent     *folder.FolderNode
	project    *project.Project
	ModuleName string
}

func NewProjectAddModule(parent *folder.FolderNode, project *project.Project, moduleName string) *ProjectAddModule {
	return &ProjectAddModule{
		parent:     parent,
		project:    project,
		ModuleName: moduleName,
	}
}

func (e *ProjectAddModule) GetModule(path *node.Path) node.Node {
	query := node.NewNodeQuery(e.ModuleName)

	module := e.project.RootNode.FindWithPath(query, path)

	if module == nil {
		return nil
	}

	return module
}

func (a *ProjectAddModule) Run() error {
	module := folder.NewModuleNode(a.ModuleName)
	a.parent.AddElement(module)

	return nil
}

func (a *ProjectAddModule) Description() string {
	return "Many application are organized into modules, which are used to group related functionality together. This action adds a new module to the project." +
		" When the developer wants to add a new module, for new group of functionality, or to organize domain logic, we can use this action."
}

func (a *ProjectAddModule) Factory() Action {
	return NewActionFactory[*ProjectAddModule]().
		Question("name", "What is the name of the module you want to add?").
		Question("path", "What is the path where the module should be added?").
		Build(func(answers Answers) *ProjectAddModule {
			return NewProjectAddModule(a.project.Root().FindChildFolder(node.NewNodeQuery(answers.Get("path"))), a.project, answers.Get("name"))
		})
}
