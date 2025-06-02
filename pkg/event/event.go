package event

type Event interface {
	Name() string
	Description() string
}

type CommandResult Event
