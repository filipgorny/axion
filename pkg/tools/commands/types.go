package commands

import "github.com/filipgorny/axion/pkg/event"

type CommandResult interface {
	event.Event
}

type CommandParam string
