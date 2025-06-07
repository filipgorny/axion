package service

import (
	"github.com/filipgorny/axion/pkg/node"
	"github.com/filipgorny/axion/pkg/node/model"
)

type ServiceNode struct {
	path    *node.Path
	name    string
	Methods []*ServiceMethod
}

func (s *ServiceNode) Type() node.NodeType {
	return "service"
}

func (s *ServiceNode) Path() *node.Path {
	if s.path == nil {
		s.path = node.NewPath(s.name)
	}

	return s.path
}

func (s *ServiceNode) Name() string {
	return s.name
}

func NewServiceNode(name string, methods ...*ServiceMethod) *ServiceNode {
	return &ServiceNode{
		name:    name,
		Methods: append(make([]*ServiceMethod, 0), methods...),
	}
}

type ServiceMethod struct {
	Name       string
	Parameters []*ServiceParameter
	Returns    []*ServiceReturn
}

type ServiceParameter struct {
	Name   string
	Type   node.DataType
	Entity *model.EntityNode
}

type ServiceReturn struct {
	Name string
	Type node.DataType
}

func NewParameter(name string, dataType node.DataType) *ServiceParameter {
	return &ServiceParameter{
		Name: name,
		Type: dataType,
	}
}

func NewEntityParameter(name string, entity *model.EntityNode) *ServiceParameter {
	return &ServiceParameter{
		Name:   name,
		Type:   node.Entity,
		Entity: entity,
	}
}

func NewReturn(name string, dataType node.DataType) *ServiceReturn {
	return &ServiceReturn{
		Name: name,
		Type: dataType,
	}
}

func NewVoidReturn() *ServiceReturn {
	return &ServiceReturn{
		Name: "void",
		Type: node.Void,
	}
}

func NewMethod(name string, returnType *ServiceReturn, parameters ...*ServiceParameter) *ServiceMethod {
	return &ServiceMethod{
		Name:       name,
		Parameters: append(make([]*ServiceParameter, 0), parameters...),
		Returns:    append(make([]*ServiceReturn, 0), returnType),
	}
}

func (s *ServiceNode) AddMethods(methods ...*ServiceMethod) *ServiceNode {
	s.Methods = append(s.Methods, methods...)

	return s
}

func (s *ServiceNode) Characteristics() node.NodeCharacteristics {
	return []string{
		"it is used to define a service in the application",
		"it can contain methods that define the service's functionality",
		"it often groups methods that operate on other nodes or data structures",
	}
}
