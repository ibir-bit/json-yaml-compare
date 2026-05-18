package formatters

import (
	"fmt"
	"strings"
)

func FormatPlain(nodes []DiffNode, parentPath string) string {
	var lines []string

	for _, node := range nodes {
		currentPath := node.Key
		if parentPath != "" {
			currentPath = fmt.Sprintf("%s.%s", parentPath, node.Key)
		}

		switch node.Status {
		case "nested":
			nestedResult := FormatPlain(node.Children, currentPath)
			if nestedResult != "" {
				lines = append(lines, nestedResult)
			}
		case "added":
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", currentPath, stringify(node.Value)))
		case "removed":
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", currentPath))
		case "changed":
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, stringify(node.OldValue), stringify(node.NewValue)))
		case "unchanged":
			continue
		}
	}

	return strings.Join(lines, "\n")
}

func stringify(v interface{}) string {
	switch val := v.(type) {
	case map[string]interface{}:
		return "[complex value]"
	case string:
		return fmt.Sprintf("'%s'", val)
	case nil:
		return "null"
	case bool, int, float64:
		return fmt.Sprintf("%v", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
