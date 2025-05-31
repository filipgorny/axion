package dto

import (
	"github.com/filipgorny/axion/pkg/node"
)

type DtoNode struct {
	Name       string
	Properties []DtoProperty
}

func (d *DtoNode) Type() string {
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
