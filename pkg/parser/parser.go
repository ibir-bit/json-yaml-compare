package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ReadFile читает файл по пути и возвращает map[string]interface{}
func ReadFile(path string) (map[string]interface{}, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve absolute path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("cannot parse JSON: %w", err)
		}
		return result, nil
	case ".yml", ".yaml":
		var result map[string]interface{}
		if err := yaml.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("cannot parse YAML: %w", err)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}
}
