package model

import (
	"github.com/filipgorny/axion/pkg/node"
)

type EntityNode struct {
	Name       string
	Properties []*Property
	Children   []*EntityNode
}

func NewEntityNode(name string, properties []*Property, children []*EntityNode) *EntityNode {
	entity := &EntityNode{
		Name:       name,
		Properties: make([]*Property, 0),
		Children:   make([]*EntityNode, 0),
	}

	for _, prop := range properties {
		entity.Properties = append(entity.Properties, prop)
	}

	for _, child := range children {
		entity.Children = append(entity.Children, child)
	}

	return entity
}

func (e *EntityNode) Type() string {
	return e.Name
}

type Property struct {
	Name string
	Type node.DataType
}
