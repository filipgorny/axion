package event

import "github.com/filipgorny/axion/pkg/project"

type NewModuleData struct {
	name string
}

func NewModuleAdded(project project.Project) *NewModuleData {
	return &NewModuleData{
		name: project.Name,
	}
}

func (e *NewModuleData) Name() string {
	return "new module " + e.name
}

func (e *NewModuleData) Description() string {
	return "Module " + e.name + " has been created"
}
