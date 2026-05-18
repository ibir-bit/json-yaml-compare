package gendiff

import "code/pkg/gendiff/formatters"

func GenDiff(data1, data2 map[string]interface{}, formatName string) (string, error) {
	diffTree := BuildDiff(data1, data2)
	return formatters.Format(diffTree, formatName)
}
