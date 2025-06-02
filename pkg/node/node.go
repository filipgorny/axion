package node

type NodeType string

type Node interface {
	Name() string
	Type() NodeType
	Characteristics() NodeCharacteristics
	Path() *Path
}

type NodeCharacteristics []string

func Characteristics(characteristics ...string) NodeCharacteristics {
	result := make([]string, 0)

	for _, value := range characteristics {
		result = append(result, value)
	}

	return result
}
