package node

import "strings"

type Path struct {
	Elements []string
}

func NewPath(elements ...string) *Path {
	return &Path{
		Elements: elements,
	}
}

func (p *Path) String() string {
	return strings.Join(p.Elements, ".")
}

func (p *Path) Add(element string) *Path {
	return &Path{
		Elements: append(p.Elements, element),
	}
}
