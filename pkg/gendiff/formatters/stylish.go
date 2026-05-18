package formatters

import (
	"fmt"
	"sort"
	"strings"
)

func FormatStylish(nodes []DiffNode, depth int) string {
	var builder strings.Builder
	builder.WriteString("{\n")

	indentSize := depth * 4
	prefix := strings.Repeat(" ", indentSize-2)
	bracketIndent := strings.Repeat(" ", indentSize-4)

	for _, node := range nodes {
		switch node.Status {
		case "nested":
			builder.WriteString(fmt.Sprintf("%s  %s: %s\n", prefix, node.Key, FormatStylish(node.Children, depth+1)))
		case "unchanged":
			val := formatValueStylish(node.Value, depth+1)
			builder.WriteString(fmt.Sprintf("%s  %s: %s\n", prefix, node.Key, val))
		case "added":
			val := formatValueStylish(node.Value, depth+1)
			builder.WriteString(fmt.Sprintf("%s+ %s: %s\n", prefix, node.Key, val))
		case "removed":
			val := formatValueStylish(node.Value, depth+1)
			builder.WriteString(fmt.Sprintf("%s- %s: %s\n", prefix, node.Key, val))
		case "changed":
			oldVal := formatValueStylish(node.OldValue, depth+1)
			newVal := formatValueStylish(node.NewValue, depth+1)
			builder.WriteString(fmt.Sprintf("%s- %s: %s\n", prefix, node.Key, oldVal))
			builder.WriteString(fmt.Sprintf("%s+ %s: %s\n", prefix, node.Key, newVal))
		}
	}

	builder.WriteString(bracketIndent + "}")
	return builder.String()
}

func formatValueStylish(v interface{}, depth int) string {
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

		indentSize := depth * 4
		currentIndent := strings.Repeat(" ", indentSize)
		bracketIndent := strings.Repeat(" ", indentSize-4)

		var builder strings.Builder
		builder.WriteString("{\n")
		for _, k := range keys {
			builder.WriteString(fmt.Sprintf("%s%s: %s\n", currentIndent, k, formatValueStylish(val[k], depth+1)))
		}
		builder.WriteString(bracketIndent + "}")
		return builder.String()
	case nil:
		return "null"
	case string:
		return val
	default:
		return fmt.Sprintf("%v", val)
	}
}
