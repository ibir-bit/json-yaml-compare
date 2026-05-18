package gendiff

import (
	"fmt"
	"sort"
	"strings"
)

func formatStylish(nodes []DiffNode, depth int) string {
	var builder strings.Builder
	builder.WriteString("{\n")

	// Базовый отступ для текущего уровня (4 пробела на глубину)
	// Знаки +/- ставятся на 2 пробела левее основного отступа ключа
	indentSize := depth * 4
	prefix := strings.Repeat(" ", indentSize-2)
	bracketIndent := strings.Repeat(" ", indentSize-4)

	for _, node := range nodes {
		switch node.Status {
		case "nested":
			// Для вложенных узлов используем "  " (два пробела) после префикса, чтобы ключ стоял ровно
			builder.WriteString(fmt.Sprintf("%s  %s: %s\n", prefix, node.Key, formatStylish(node.Children, depth+1)))
		case "unchanged":
			builder.WriteString(fmt.Sprintf("%s  %s: %s\n", prefix, node.Key, formatValue(node.Value, depth+1)))
		case "added":
			builder.WriteString(fmt.Sprintf("%s+ %s: %s\n", prefix, node.Key, formatValue(node.Value, depth+1)))
		case "removed":
			builder.WriteString(fmt.Sprintf("%s- %s: %s\n", prefix, node.Key, formatValue(node.Value, depth+1)))
		case "changed":
			builder.WriteString(fmt.Sprintf("%s- %s: %s\n", prefix, node.Key, formatValue(node.OldValue, depth+1)))
			builder.WriteString(fmt.Sprintf("%s+ %s: %s\n", prefix, node.Key, formatValue(node.NewValue, depth+1)))
		}
	}

	builder.WriteString(bracketIndent + "}")
	return builder.String()
}

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

		indentSize := depth * 4
		currentIndent := strings.Repeat(" ", indentSize)
		bracketIndent := strings.Repeat(" ", indentSize-4)

		var builder strings.Builder
		builder.WriteString("{\n")
		for _, k := range keys {
			builder.WriteString(fmt.Sprintf("%s%s: %s\n", currentIndent, k, formatValue(val[k], depth+1)))
		}
		builder.WriteString(bracketIndent + "}")
		return builder.String()

	case nil:
		// ИЗМЕНЕНИЕ 1: Возвращаем <nil>, как просит тест
		return "<nil>"

	case string:
		// ИЗМЕНЕНИЕ 2: Чтобы избежать проблем с невидимыми пробелами,
		// убеждаемся, что пустая строка выводится без лишних символов.
		if val == "" {
			return ""
		}
		return val

	default:
		return fmt.Sprintf("%v", val)
	}
}
