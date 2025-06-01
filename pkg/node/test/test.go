package test

import (
	"github.com/filipgorny/axion/pkg/node"
)

type TestNode struct {
	Name string
}

func NewTestNode(name string) *TestNode {
	return &TestNode{
		Name: name,
	}
}

func (t *TestNode) Type() node.NodeType {
	return "test"
}

func (t *TestNode) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"it is used to represent a single test case",
		"it can be part of a test suite",
		"it is often used to verify the functionality of a component or feature",
	}
}
