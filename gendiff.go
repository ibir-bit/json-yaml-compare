package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type DiffNode struct {
	Key      string
	Status   string
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
	Children []DiffNode
}

func GenDiff(path1, path2 string, format string) (string, error) {
	data1, err := readAndParse(path1)
	if err != nil {
		return "", err
	}

	data2, err := readAndParse(path2)
	if err != nil {
		return "", err
	}

	switch format {
	case "plain":
		return FormatPlain(data1, data2), nil
	case "json":
		return FormatJSON(data1, data2), nil
	case "stylish":
		fallthrough
	default:
		return FormatStylish(data1, data2), nil
	}
}

func readAndParse(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))
	result := make(map[string]interface{})

	if ext == ".json" {
		err = json.Unmarshal(content, &result)
	} else {
		err = yaml.Unmarshal(content, &result)
	}
	return result, err
}

func isNil(v interface{}) bool {
	return v == nil
}

func formatLine(prefix, key, value string) string {
	if value == "" {
		return fmt.Sprintf("%s%s:\n", prefix, key)
	}
	return fmt.Sprintf("%s%s: %s\n", prefix, key, value)
}

func FormatStylish(data1, data2 map[string]interface{}) string {
	diffTree := buildDiff(data1, data2)
	return formatStylishRecursive(diffTree, 1)
}

func formatStylishRecursive(nodes []DiffNode, depth int) string {
	var builder strings.Builder
	builder.WriteString("{\n")

	indent := strings.Repeat("    ", depth)
	signIndent := strings.Repeat("    ", depth-1) + "  "

	for _, node := range nodes {
		switch node.Status {
		case "nested":
			builder.WriteString(formatLine(indent, node.Key, formatStylishRecursive(node.Children, depth+1)))
		case "unchanged":
			builder.WriteString(formatLine(indent, node.Key, formatValueStylish(node.Value, depth+1)))
		case "added":
			builder.WriteString(formatLine(signIndent+"+ ", node.Key, formatValueStylish(node.Value, depth+1)))
		case "removed":
			builder.WriteString(formatLine(signIndent+"- ", node.Key, formatValueStylish(node.Value, depth+1)))
		case "changed":
			builder.WriteString(formatLine(signIndent+"- ", node.Key, formatValueStylish(node.OldValue, depth+1)))
			builder.WriteString(formatLine(signIndent+"+ ", node.Key, formatValueStylish(node.NewValue, depth+1)))
		}
	}
	builder.WriteString(strings.Repeat("    ", depth-1) + "}")
	return builder.String()
}

func formatValueStylish(v interface{}, depth int) string {
	if isNil(v) {
		return "null"
	}
	switch val := v.(type) {
	case bool, int, float64:
		return fmt.Sprintf("%v", val)
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
			// Используем ту же функцию сборки строки, чтобы убрать проблему и во вложенных словарях
			builder.WriteString(formatLine(indent, k, formatValueStylish(val[k], depth+1)))
		}
		builder.WriteString(strings.Repeat("    ", depth-1) + "}")
		return builder.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}

func FormatPlain(data1, data2 map[string]interface{}) string {
	diffTree := buildDiff(data1, data2)
	lines := buildPlainLines(diffTree, "")
	return strings.Join(lines, "\n")
}

func buildPlainLines(nodes []DiffNode, parentPath string) []string {
	var lines []string

	for _, node := range nodes {
		currentPath := node.Key
		if parentPath != "" {
			currentPath = parentPath + "." + node.Key
		}

		switch node.Status {
		case "nested":
			lines = append(lines, buildPlainLines(node.Children, currentPath)...)
		case "added":
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", currentPath, formatValuePlain(node.Value)))
		case "removed":
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", currentPath))
		case "changed":
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, formatValuePlain(node.OldValue), formatValuePlain(node.NewValue)))
		case "unchanged":
			// Plain игнорирует неизмененные элементы
		}
	}
	return lines
}

func formatValuePlain(v interface{}) string {
	if isNil(v) {
		return "null"
	}
	switch val := v.(type) {
	case map[string]interface{}:
		return "[complex value]"
	case string:
		return fmt.Sprintf("'%s'", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func treeToMap(nodes []DiffNode) map[string]interface{} {
	result := make(map[string]interface{})
	for _, node := range nodes {
		nodeData := make(map[string]interface{})
		nodeData["status"] = node.Status

		switch node.Status {
		case "nested":
			nodeData["children"] = treeToMap(node.Children)
		case "changed":
			nodeData["oldValue"] = node.OldValue
			nodeData["newValue"] = node.NewValue
		default:
			nodeData["value"] = node.Value
		}

		result[node.Key] = nodeData
	}
	return result
}

func FormatJSON(data1, data2 map[string]interface{}) string {
	diffTree := buildDiff(data1, data2)
	mappedTree := treeToMap(diffTree)

	bytes, err := json.Marshal(mappedTree)
	if err != nil {
		return "{}"
	}
	return string(bytes)
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
			} else if reflect.DeepEqual(v1, v2) {
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
