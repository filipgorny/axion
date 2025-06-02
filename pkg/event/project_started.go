package event

import "github.com/filipgorny/axion/pkg/project"

type ProjectStarted struct {
	project *project.Project
}

func NewProjectStarted(project project.Project) *ProjectStarted {
	return &ProjectStarted{
		project: &project,
	}
}

func (e *ProjectStarted) Name() string {
	return "project " + e.project.Name + " started"
}

func (e *ProjectStarted) Description() string {
	return "Project " + e.project.Name + " has been started"
}
