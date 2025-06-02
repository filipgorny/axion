package engine

type Engine interface {
	Ask(question string) (string, error)
}
