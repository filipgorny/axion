package commands

import (
	"github.com/filipgorny/axion/pkg/project"
)

type Command interface {
	Execute(project project.Project) func(params ...CommandParam) CommandResult
	Description() string
}
