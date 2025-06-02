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

func NewEntity(name string) *EntityNode {
	entity := &EntityNode{
		Name:       name,
		Properties: make([]*Property, 0),
		Children:   make([]*EntityNode, 0),
	}

	return entity
}

func (e *EntityNode) Type() string {
	return "entity"
}

type Property struct {
	Name string
	Type node.DataType
}

func NewProperty(name string, dataType node.DataType) *Property {
	return &Property{
		Name: name,
		Type: dataType,
	}
}

func (e *EntityNode) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"it is used to represent a domain entity",
		"it can contain properties that define the entity's state in the runtime",
		"it can have child entities that represent relationships or nested structures",
		"it can correspond to a table in a database",
	}
}
