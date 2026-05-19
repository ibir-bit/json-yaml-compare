package code

import (
	"code/pkg/gendiff"
	"os"

	"gopkg.in/yaml.v3"
)

// GenDiff — это функция, которую вызывают тесты Хекслета
func GenDiff(path1, path2 string, format string) (string, error) {
	_ = format // пока не используем

	data1, err := readAndParse(path1)
	if err != nil {
		return "", err
	}

	data2, err := readAndParse(path2)
	if err != nil {
		return "", err
	}

	return gendiff.GenDiffRecursive(data1, data2), nil
}

// Перенесли сюда, чтобы линтер видел её из пакета code
func readAndParse(filepath string) (map[string]interface{}, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
