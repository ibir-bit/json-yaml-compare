package formatters

import "fmt"

type DiffNode struct {
	Key      string
	Status   string // "nested", "added", "removed", "unchanged", "changed"
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
	Children []DiffNode
}

func Format(nodes []DiffNode, formatName string) (string, error) {
	switch formatName {
	case "stylish":
		return FormatStylish(nodes, 1), nil
	case "plain":
		return FormatPlain(nodes, ""), nil
	default:
		return "", fmt.Errorf("unknown format: %s", formatName)
	}
}
