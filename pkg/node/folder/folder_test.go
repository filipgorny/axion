package folder

import (
	"testing"

	"github.com/filipgorny/axion/pkg/node"
)

func TestFindingElement(t *testing.T) {
	// Create a new folder
	testFolder := NewFolderNode("testFolder", node.NewDummyNode(), node.NewDummyNode(), node.NewDummyNode())

	// Add some elements to the folder
	testFolder.AddElement(node.NewStandardNode("element1", "type1", nil, nil))
	testFolder.AddElement(node.NewDummyNode())

	// Test finding an element by name
	element := testFolder.Find(node.NewNodeQuery("element1"))
	if element == nil || element.Name() != "element1" {
		t.Errorf("Expected to find element 'element1', but got %v", element)
	}

	// Test finding an element that does not exist
	element = testFolder.Find(node.NewNodeQuery("nonexistent"))
	if element != nil {
		t.Errorf("Expected to not find any element, but got %v", element)
	}
}
