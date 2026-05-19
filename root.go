package code

import (
	"code/pkg/gendiff" // Убедитесь, что здесь ваш правильный импорт пакета
)

// GenDiff читает два файла, сравнивает их и возвращает результат в виде строки.
// Именно эту функцию ищут тесты Хекслета.
// Добавьте третий параметр 'format'
func GenDiff(path1, path2 string, format string) (string, error) {
	// Временно игнорируем формат, если он еще не реализован в логике
	_ = format

	// 1. Читаем и парсим файлы (код, который вы уже написали)
	data1, err := readAndParse(path1)
	if err != nil {
		return "", err
	}

	data2, err := readAndParse(path2)
	if err != nil {
		return "", err
	}

	// 2. Вызываем основную логику
	diff := gendiff.GenDiffRecursive(data1, data2)

	return diff, nil
}
