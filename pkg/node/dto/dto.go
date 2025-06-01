package dto

import (
	"github.com/filipgorny/axion/pkg/node"
)

type DtoNode struct {
	Name       string
	Properties []DtoProperty
}

func NewDtoNode(name string, properties ...DtoProperty) *DtoNode {
	return &DtoNode{
		Name:       name,
		Properties: append(make([]DtoProperty, 0), properties...),
	}
}

func (d *DtoNode) Type() node.NodeType {
	return "dto"
}

type DtoProperty interface {
	Name() string
	Type() node.DataType
}

type DtoPropertyPrimitive struct {
	Name string
}

type DtoPropertyCollection struct {
	Name     string
	Children []DtoProperty
}

func (d *DtoPropertyPrimitive) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"doesn't include ID",
		"it is often use to represent data structures passed to the methods",
	}
}
