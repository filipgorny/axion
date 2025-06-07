package action

type Action interface {
	Run() error
	Description() string
	Factory() Action
}
