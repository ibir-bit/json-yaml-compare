package gendiff

import (
	"fmt"
	"sort"
	"strings"
)

// formatStylish форматирует diff в красивый стиль stylish
func formatStylish(nodes []DiffNode, depth int) string {
	indent := func(level int) string {
		return strings.Repeat("  ", level)
	}

	var formatValue func(v interface{}, depth int) string
	formatValue = func(v interface{}, depth int) string {
		switch val := v.(type) {
		case map[string]interface{}:
			if len(val) == 0 {
				return "{}"
			}
			keys := make([]string, 0, len(val))
			for k := range val {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			result := "{\n"
			for _, k := range keys {
				result += fmt.Sprintf("%s%s: %s\n", indent(depth), k, formatValue(val[k], depth+1))
			}
			result += indent(depth-1) + "}"
			return result
		case nil:
			return "<nil>"
		default:
			return fmt.Sprintf("%v", val)
		}
	}

	result := "{\n"
	for _, node := range nodes {
		switch node.Status {
		case "nested":
			result += fmt.Sprintf("%s%s: %s", indent(depth), node.Key, formatStylish(node.Children, depth+1))
		case "unchanged":
			result += fmt.Sprintf("%s%s: %s\n", indent(depth), node.Key, formatValue(node.Value, depth+1))
		case "added":
			result += fmt.Sprintf("%s+ %s: %s\n", indent(depth-1), node.Key, formatValue(node.Value, depth))
		case "removed":
			result += fmt.Sprintf("%s- %s: %s\n", indent(depth-1), node.Key, formatValue(node.Value, depth))
		case "changed":
			result += fmt.Sprintf("%s- %s: %s\n", indent(depth-1), node.Key, formatValue(node.OldValue, depth))
			result += fmt.Sprintf("%s+ %s: %s\n", indent(depth-1), node.Key, formatValue(node.NewValue, depth))
		}
	}
	result += indent(depth-1) + "}"
	return result
}
