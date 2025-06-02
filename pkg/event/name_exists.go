package event

type NameExists struct {
	name       string
	kindOfData string
}

func NewNameExists(name string, kindOfData string) *NameExists {
	return &NameExists{
		name:       name,
		kindOfData: kindOfData,
	}
}

func (e *NameExists) Name() string {
	return "name " + e.name + " already exists"
}

func (e *NameExists) Description() string {
	return e.kindOfData + " with a name " + e.name + " already exists in the project"
}

func NewInvalidNameError(name string, kindOfData string) *NameExists {
	return &NameExists{
		name:       name,
		kindOfData: kindOfData,
	}
}
