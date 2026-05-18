package gendiff

import (
	"code/pkg/gendiff/formatters"
	"code/pkg/parser" // Импортируем ваш парсер
)

// GenDiff теперь принимает пути к файлам (string), а не данные (map)
func GenDiff(path1, path2 string, formatName string) (string, error) {
	// 1. Внутри функции сами читаем и парсим файлы
	data1, err := parser.ReadFile(path1)
	if err != nil {
		return "", err
	}

	data2, err := parser.ReadFile(path2)
	if err != nil {
		return "", err
	}

	// 2. Строим дерево различий
	diffTree := BuildDiff(data1, data2)

	// 3. Форматируем результат
	return formatters.Format(diffTree, formatName)
}
