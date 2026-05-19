package code

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"code/pkg/gendiff"

	"gopkg.in/yaml.v3"
)

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
		return gendiff.FormatPlain(data1, data2), nil
	case "json":
		return gendiff.FormatJSON(data1, data2), nil
	case "stylish":
		fallthrough
	default:
		return gendiff.FormatStylish(data1, data2), nil
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
