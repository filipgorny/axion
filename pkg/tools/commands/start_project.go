package commands

import (
	"github.com/filipgorny/axion/pkg/event"
	"github.com/filipgorny/axion/pkg/project"
)

type StartProject struct {
}

func (c *StartProject) Execute() func(params ...CommandParam) CommandResult {
	return func(params ...CommandParam) CommandResult {
		name := string(params[0])

		project := project.NewProject(name)

		return event.NewProjectStarted(*project)
	}
}
