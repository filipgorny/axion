package node

import (
	"strings"
)

type NodeType string

type Node interface {
	Type() NodeType
	Characteristics() []string
}

type NodeCharacteristics []string

func Characeteristics(characteristics string) NodeCharacteristics {
	result := make([]string, 0)

	characteristicsList := strings.Split(characteristics, "\n")

	for _, value := range characteristicsList {
		result = append(result, value)
	}

	return result
}
