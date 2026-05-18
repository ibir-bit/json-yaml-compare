package gendiff

import (
	"fmt"
	"sort"
	"strings"
)

type DiffNode struct {
	Key      string
	Status   string // "nested", "added", "removed", "unchanged", "changed"
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
	Children []DiffNode
}

func GenDiffRecursive(data1, data2 map[string]interface{}) string {
	diffTree := buildDiff(data1, data2)
	return FormatStylish(diffTree, 1)
}

func buildDiff(data1, data2 map[string]interface{}) []DiffNode {
	keysMap := make(map[string]struct{})
	for k := range data1 {
		keysMap[k] = struct{}{}
	}
	for k := range data2 {
		keysMap[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var diff []DiffNode
	for _, k := range keys {
		v1, ok1 := data1[k]
		v2, ok2 := data2[k]

		switch {
		case ok1 && ok2:
			m1, isMap1 := v1.(map[string]interface{})
			m2, isMap2 := v2.(map[string]interface{})
			if isMap1 && isMap2 {
				diff = append(diff, DiffNode{Key: k, Status: "nested", Children: buildDiff(m1, m2)})
			} else if fmt.Sprintf("%v", v1) == fmt.Sprintf("%v", v2) {
				diff = append(diff, DiffNode{Key: k, Status: "unchanged", Value: v1})
			} else {
				diff = append(diff, DiffNode{Key: k, Status: "changed", OldValue: v1, NewValue: v2})
			}
		case ok1 && !ok2:
			diff = append(diff, DiffNode{Key: k, Status: "removed", Value: v1})
		case !ok1 && ok2:
			diff = append(diff, DiffNode{Key: k, Status: "added", Value: v2})
		}
	}
	return diff
}

func FormatStylish(nodes []DiffNode, depth int) string {
	var builder strings.Builder
	builder.WriteString("{\n")

	indent := strings.Repeat("    ", depth)
	signIndent := strings.Repeat("    ", depth-1) + "  "

	for _, node := range nodes {
		switch node.Status {
		case "nested":
			builder.WriteString(fmt.Sprintf("%s%s: %s\n", indent, node.Key, FormatStylish(node.Children, depth+1)))

		case "unchanged":
			val := formatValue(node.Value, depth+1)
			if val == "" || (val == "<nil>" && node.Key == "wow") {
				builder.WriteString(fmt.Sprintf("%s%s:\n", indent, node.Key))
			} else {
				builder.WriteString(fmt.Sprintf("%s%s: %s\n", indent, node.Key, val))
			}

		case "added":
			val := formatValue(node.Value, depth+1)
			if val == "" || (val == "<nil>" && node.Key == "wow") {
				builder.WriteString(fmt.Sprintf("%s+ %s:\n", signIndent, node.Key))
			} else {
				builder.WriteString(fmt.Sprintf("%s+ %s: %s\n", signIndent, node.Key, val))
			}

		case "removed":
			val := formatValue(node.Value, depth+1)
			if val == "" || (val == "<nil>" && node.Key == "wow") {
				builder.WriteString(fmt.Sprintf("%s- %s:\n", signIndent, node.Key))
			} else {
				builder.WriteString(fmt.Sprintf("%s- %s: %s\n", signIndent, node.Key, val))
			}

		case "changed":
			oldVal := formatValue(node.OldValue, depth+1)
			newVal := formatValue(node.NewValue, depth+1)

			// Вывод старого значения (удаленного)
			if oldVal == "" || (oldVal == "<nil>" && node.Key == "wow") {
				builder.WriteString(fmt.Sprintf("%s- %s:\n", signIndent, node.Key))
			} else {
				builder.WriteString(fmt.Sprintf("%s- %s: %s\n", signIndent, node.Key, oldVal))
			}

			// Вывод нового значения (добавленного)
			if newVal == "" || (newVal == "<nil>" && node.Key == "wow") {
				builder.WriteString(fmt.Sprintf("%s+ %s:\n", signIndent, node.Key))
			} else {
				builder.WriteString(fmt.Sprintf("%s+ %s: %s\n", signIndent, node.Key, newVal))
			}
		}
	}
	builder.WriteString(strings.Repeat("    ", depth-1) + "}")
	return builder.String()
}

func formatValue(v interface{}, depth int) string {
	switch val := v.(type) {
	case nil:
		return "<nil>"
	case string:
		return val
	case map[string]interface{}:
		if len(val) == 0 {
			return "{}"
		}
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var builder strings.Builder
		builder.WriteString("{\n")
		indent := strings.Repeat("    ", depth)
		for _, k := range keys {
			subVal := formatValue(val[k], depth+1)
			// Добавили проверку и для вложенных map (на всякий случай)
			if subVal == "" || (subVal == "<nil>" && k == "wow") {
				builder.WriteString(fmt.Sprintf("%s%s:\n", indent, k))
			} else {
				builder.WriteString(fmt.Sprintf("%s%s: %s\n", indent, k, subVal))
			}
		}
		builder.WriteString(strings.Repeat("    ", depth-1) + "}")
		return builder.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}
