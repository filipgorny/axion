package node

type NodeQuery struct {
	Name string
	Type NodeType
	Path *Path
}

func NewNodeQuery(name string) *NodeQuery {
	return &NodeQuery{
		Name: name,
	}
}

func NewDetailedNodeQuery(name string, nodeType NodeType, path *Path) *NodeQuery {
	return &NodeQuery{
		Name: name,
		Type: nodeType,
		Path: path,
	}
}

func (q *NodeQuery) Matches(node Node) bool {
	if q.Name != "" && node.Name() != q.Name {
		return false
	}
	if q.Type != "" && node.Type() != q.Type {
		return false
	}
	if q.Path != nil && node.Path() != nil && node.Path().String() != q.Path.String() {
		return false
	}
	return true
}
