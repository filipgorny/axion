package test

import (
	"github.com/filipgorny/axion/pkg/node"
)

type TestSuiteNode struct {
	Name     string
	Children []*TestNode
}

func NewTestSuiteNode(name string) *TestSuiteNode {
	return &TestSuiteNode{
		Name:     name,
		Children: make([]*TestNode, 0),
	}
}

func (t *TestSuiteNode) Type() node.NodeType {
	return "test_suite"
}

func (t *TestSuiteNode) AddChild(child *TestNode) {
	t.Children = append(t.Children, child)
}

func (t *TestSuiteNode) RemoveChild(child *TestNode) {
	for i, c := range t.Children {
		if c == child {
			t.Children = append(t.Children[:i], t.Children[i+1:]...)
			return
		}
	}
}

func (t *TestSuiteNode) GetChildren() []*TestNode {
	return t.Children
}

func (t *TestSuiteNode) Characteristics() node.NodeCharacteristics {
	return node.NodeCharacteristics{
		"it is used to group tests together in a suite",
		"it can contain multiple test nodes",
		"it is often used to represent a collection of related tests",
	}
}
