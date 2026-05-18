package gendiff

import (
	"fmt"
	"strings"
)

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case map[string]interface{}:
		if len(v) == 0 {
			return "{}"
		}
		result := "{"
		for k, val := range v {
			result += fmt.Sprintf(" %s: %s", k, formatValue(val))
		}
		return result + " }"
	default:
		if v == nil {
			return "null"
		}
		return fmt.Sprintf("%v", v)
	}
}

func formatStylish(nodes []DiffNode, depth int) string {
	indent := func(level int) string {
		return strings.Repeat("    ", level)
	}

	result := "{\n"
	for _, node := range nodes {
		switch node.Status {
		case "nested":
			result += fmt.Sprintf("%s%s: %s", indent(depth), node.Key, formatStylish(node.Children, depth+1))
		case "unchanged":
			result += fmt.Sprintf("%s%s: %s\n", indent(depth), node.Key, formatValue(node.Value))
		case "added":
			result += fmt.Sprintf("%s+ %s: %s\n", indent(depth), node.Key, formatValue(node.Value))
		case "removed":
			result += fmt.Sprintf("%s- %s: %s\n", indent(depth), node.Key, formatValue(node.Value))
		case "changed":
			result += fmt.Sprintf("%s- %s: %s\n", indent(depth), node.Key, formatValue(node.OldValue))
			result += fmt.Sprintf("%s+ %s: %s\n", indent(depth), node.Key, formatValue(node.NewValue))
		}
	}
	result += indent(depth-1) + "}"
	return result
}
