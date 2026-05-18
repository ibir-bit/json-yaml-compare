package gendiff

import (
	"code/pkg/gendiff/formatters"
	"code/pkg/parser"
)

// GenDiff теперь принимает пути к файлам и парсит их сам
func GenDiff(path1, path2 string, formatName string) (string, error) {
	data1, err := parser.ReadFile(path1)
	if err != nil {
		return "", err
	}

	data2, err := parser.ReadFile(path2)
	if err != nil {
		return "", err
	}

	diffTree := BuildDiff(data1, data2)
	return formatters.Format(diffTree, formatName)
}
