package parser

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/filipgorny/axion/pkg/node/service"
)

func TestParser(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "/home/filip/Projects/filipgorny/axion/pkg/"

	// Create a sample project structure
	projectName := "node"
	projectPath := filepath.Join(tempDir, projectName)

	parser := NewParser()

	project := parser.Parse(projectPath)

	for _, child := range project.RootNode.Children {
		log.Printf("Node: %s, Type: %s", child.Name(), child.Type())

		log.Printf("Entering node: %s", child.Path().String())

		if child.Type() == "service" {
			serviceNode := child.(*service.ServiceNode)
			log.Printf("Service: %s", serviceNode.Name())
			for _, method := range serviceNode.Methods {
				log.Printf("  Method: %s", method.Name)
				for _, param := range method.Parameters {
					log.Printf("Parameter: %s", param.Name)
				}
				for _, ret := range method.Returns {
					log.Printf("Return: %s", ret.Name)
				}
			}
		}
	}
}
