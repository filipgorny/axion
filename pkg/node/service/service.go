package service

type ServiceNode struct {
	Name    string
	Methods []ServiceMethod
}

func NewServiceNode(name string, methods ...ServiceMethod) *ServiceNode {
	return &ServiceNode{
		Name:    name,
		Methods: append(make([]ServiceMethod, 0), methods...),
	}
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

func (s *ServiceNode) Characteristics() []string {
	return []string{
		"it is used to define a service in the application",
		"it can contain methods that define the service's functionality",
		"it often groups methods that operate on other nodes or data structures",
	}
}
