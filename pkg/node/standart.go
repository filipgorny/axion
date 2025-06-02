package node

type StandardNode struct {
	name     string
	typeName NodeType
	path     *Path
}

func NewStandardNode(name string, nodeType NodeType, characteristics NodeCharacteristics, path *Path) *StandardNode {
	return &StandardNode{
		name:     name,
		typeName: nodeType,
	}
}

func NewDummyNode() *StandardNode {
	return &StandardNode{
		name:     "dummy",
		typeName: "dummy",
		path:     NewPath("dummy"),
	}
}

func (n *StandardNode) Name() string {
	return n.name
}

func (n *StandardNode) Type() NodeType {
	return n.typeName
}

func (n *StandardNode) Characteristics() NodeCharacteristics {
	return Characteristics(
		"it is a standard node",
		"it can be used to represent any node type",
		"it has a name, type, and path",
	)
}

func (n *StandardNode) Path() *Path {
	return n.path
}

func (n *StandardNode) SetPath(path *Path) {
	n.path = path
}

func (n *StandardNode) SetName(name string) {
	n.name = name
}

func (n *StandardNode) SetType(nodeType NodeType) {
	n.typeName = nodeType
}
