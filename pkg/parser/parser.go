package parser

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/filipgorny/axion/pkg/node/folder"
	"github.com/filipgorny/axion/pkg/node/service"
	"github.com/filipgorny/axion/pkg/project"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func ReadFile(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	body := strings.Builder{}

	for scanner.Scan() {
		body.WriteString(scanner.Text() + " ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return body.String()
}

func (p *Parser) Parse(folderName string) *project.Project {
	project := project.NewProject(folderName)

	filepath.WalkDir(folderName, func(folderName string, d fs.DirEntry, err error) error {
		log.Printf("Walking through directory: %s", folderName)

		if err != nil {
			log.Fatalf("impossible to walk directory %s: %s", folderName, err)
		}

		if d.IsDir() {
			folderDir := folder.NewModuleNode(folderName)
			project.AddModule(folderDir)

			files, err := os.ReadDir(folderName)

			if err != nil {
				log.Fatalf("impossible to read directory %s: %s", folderName, err)
			}

			for _, file := range files {
				filePath := filepath.Join(folderName, file.Name())

				if file.IsDir() {
					continue
				}

				content := ReadFile(filePath)
				body := string(content)

				if err != nil {
					log.Fatalf("impossible to read file %s: %s", file.Name(), err)
				}

				r := regexp.MustCompile(`type \w+ struct \{`)
				matches := r.FindAllString(body, -1)

				for _, match := range matches {
					log.Printf("Found match in file %s: %s", file.Name(), match)

					service := service.NewServiceNode(string(match[1]))

					if folderDir != nil {
						folderDir.AddElement(service)
					} else {
						folderDir = folder.NewModuleNode(folderName)
						folderDir.AddElement(service)
					}
				}
			}
		}

		return nil
	})

	return project
}
