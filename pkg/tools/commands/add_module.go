package commands

import (
	"github.com/filipgorny/axion/pkg/event"
	"github.com/filipgorny/axion/pkg/node/folder"
	"github.com/filipgorny/axion/pkg/project"
)

type AddModule struct {
	project project.Project
}

func (c *AddModule) Name() string {
	return "add module"
}

func NewAddModule(project project.Project, name string) *AddModule {
	return &AddModule{
		project: project,
	}
}

func (c *AddModule) Execute() func(params ...CommandParam) CommandResult {
	return func(params ...CommandParam) CommandResult {
		moduleName := string(params[0])

		moduleInstance := c.project.GetModule(moduleName)
		if moduleInstance != nil {
			return event.NewInvalidNameError(moduleName, "module")
		}

		moduleProjectInstance := folder.NewModuleNode(moduleName)
		c.project.AddModule(moduleProjectInstance)

		return event.NewModuleAdded(c.project)
	}
}

func (c *AddModule) Description() string {
	return "Lets you add a new module to the project"
}
