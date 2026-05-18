package gendiff

import (
	"fmt"
	"sort"
	"strings"
)

func formatStylish(nodes []DiffNode, depth int) string {
	indent := strings.Repeat("    ", depth-1)

	var formatValue func(interface{}, int) string
	formatValue = func(val interface{}, lvl int) string {
		switch v := val.(type) {
		case map[string]interface{}:
			if len(v) == 0 {
				return "{}"
			}

			lines := []string{"{"}
			keys := make([]string, 0, len(v))
			for k := range v {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, k := range keys {
				lines = append(lines, fmt.Sprintf(
					"%s    %s: %s",
					strings.Repeat("    ", lvl),
					k,
					formatValue(v[k], lvl+1),
				))
			}

			lines = append(lines, fmt.Sprintf("%s}", strings.Repeat("    ", lvl-1)))
			return strings.Join(lines, "\n")
		case nil:
			return "<nil>"
		default:
			return fmt.Sprintf("%v", v)
		}
	}

	lines := []string{"{"}

	for _, node := range nodes {
		switch node.Status {
		case "nested":
			lines = append(lines, fmt.Sprintf("%s    %s: %s",
				indent, node.Key, formatStylish(node.Children, depth+1),
			))
		case "unchanged":
			lines = append(lines, fmt.Sprintf("%s    %s: %s",
				indent, node.Key, formatValue(node.Value, depth+1),
			))
		case "added":
			lines = append(lines, fmt.Sprintf("%s  + %s: %s",
				indent, node.Key, formatValue(node.Value, depth+1),
			))
		case "removed":
			lines = append(lines, fmt.Sprintf("%s  - %s: %s",
				indent, node.Key, formatValue(node.Value, depth+1),
			))
		case "changed":
			lines = append(lines, fmt.Sprintf("%s  - %s: %s",
				indent, node.Key, formatValue(node.OldValue, depth+1),
			))
			lines = append(lines, fmt.Sprintf("%s  + %s: %s",
				indent, node.Key, formatValue(node.NewValue, depth+1),
			))
		}
	}

	lines = append(lines, fmt.Sprintf("%s}", strings.Repeat("    ", depth-1)))
	return strings.Join(lines, "\n")
}
