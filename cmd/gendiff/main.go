package gendiff

import (
	"fmt"
	"sort"
	"strings"
)

type DiffNode struct {
	Key      string
	Status   string      // "nested", "added", "removed", "unchanged", "changed"
	Value    interface{} // для "added" и "unchanged"
	OldValue interface{} // для "changed" и "removed"
	NewValue interface{} // для "changed" и "added"
	Children []DiffNode  // для "nested"
}

// formatStylish возвращает diff в стиле stylish
func formatStylish(nodes []DiffNode, depth int) string {
	indent := strings.Repeat("    ", depth-1) // 4 пробела на уровень

	var builder strings.Builder
	builder.WriteString("{\n")

	for _, node := range nodes {
		switch node.Status {
		case "nested":
			builder.WriteString(fmt.Sprintf("%s    %s: %s\n", indent, node.Key, formatStylish(node.Children, depth+1)))
		case "unchanged":
			builder.WriteString(fmt.Sprintf("%s    %s: %s\n", indent, node.Key, formatValue(node.Value, depth+1)))
		case "added":
			builder.WriteString(fmt.Sprintf("%s  + %s: %s\n", indent, node.Key, formatValue(node.Value, depth+1)))
		case "removed":
			builder.WriteString(fmt.Sprintf("%s  - %s: %s\n", indent, node.Key, formatValue(node.Value, depth+1)))
		case "changed":
			builder.WriteString(fmt.Sprintf("%s  - %s: %s\n", indent, node.Key, formatValue(node.OldValue, depth+1)))
			builder.WriteString(fmt.Sprintf("%s  + %s: %s\n", indent, node.Key, formatValue(node.NewValue, depth+1)))
		}
	}

	builder.WriteString(indent + "}") // закрывающая скобка на уровне открытия
	return builder.String()
}

// formatValue форматирует значение (включая map) с правильными отступами
func formatValue(v interface{}, depth int) string {
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

		indent := strings.Repeat("    ", depth)
		var builder strings.Builder
		builder.WriteString("{\n")
		for _, k := range keys {
			builder.WriteString(fmt.Sprintf("%s    %s: %s\n", indent, k, formatValue(val[k], depth+1)))
		}
		builder.WriteString(indent + "}")
		return builder.String()
	case nil:
		return "<nil>"
	default:
		return fmt.Sprintf("%v", val)
	}
}
