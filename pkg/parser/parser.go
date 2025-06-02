package parser

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/filipgorny/axion/pkg/project"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(folderName string) {
	project := project.NewProject(folderName)

	err := filepath.WalkDir(folderName, func(folderName string, d fs.DirEntry, err error) error {

		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories: %s", err)
	}

}
