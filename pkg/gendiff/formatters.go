package gendiff

import (
	"fmt"
	"strings"
)

// formatStylish форматирует diff в красивый стиль stylish
func formatStylish(nodes []DiffNode, depth int) string {
	indent := func(depth int) string {
		return strings.Repeat("  ", depth)
	}

	result := "{\n"
	for _, node := range nodes {
		switch node.Status {
		case "nested":
			result += fmt.Sprintf("%s  %s: %s", indent(depth-1), node.Key, formatStylish(node.Children, depth+1))
		case "unchanged":
			result += fmt.Sprintf("%s  %s: %v\n", indent(depth-1), node.Key, node.Value)
		case "added":
			result += fmt.Sprintf("%s+ %s: %v\n", indent(depth-1), node.Key, node.Value)
		case "removed":
			result += fmt.Sprintf("%s- %s: %v\n", indent(depth-1), node.Key, node.Value)
		case "changed":
			result += fmt.Sprintf("%s- %s: %v\n", indent(depth-1), node.Key, node.OldValue)
			result += fmt.Sprintf("%s+ %s: %v\n", indent(depth-1), node.Key, node.NewValue)
		}
	}
	result += indent(depth-1) + "}\n"
	return result
}
