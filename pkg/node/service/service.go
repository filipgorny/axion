package service

type ServiceNode struct {
	Name    string
	Methods []ServiceMethod
}

func (s *ServiceNode) Type() string {
	return "service"
}

type ServiceMethod struct {
	Name       string
	Parameters []ServiceParameter
	Returns    []ServiceReturn
}

type ServiceParameter struct {
	Name string
	Type string
}

type ServiceReturn struct {
	Name string
	Type string
}
