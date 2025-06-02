package event

type NothingHappend struct {
}

func (e *NothingHappend) Name() string {
	return "nothing happend"
}

func (e *NothingHappend) Parameters() []string {
	return []string{}
}

func NewNothingHappend() *NothingHappend {
	return &NothingHappend{}
}
