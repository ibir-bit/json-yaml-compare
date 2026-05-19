package code

import (
	"os"

	"code/pkg/gendiff" // Убедитесь, что здесь ваш правильный импорт пакета

	"gopkg.in/yaml.v3"
)

// GenDiff читает два файла, сравнивает их и возвращает результат в виде строки.
// Именно эту функцию ищут тесты Хекслета.
func GenDiff(path1, path2 string) (string, error) {
	// 1. Читаем и парсим первый файл
	data1, err := readAndParse(path1)
	if err != nil {
		return "", err
	}

	// 2. Читаем и парсим второй файл
	data2, err := readAndParse(path2)
	if err != nil {
		return "", err
	}

	// 3. Вызываем вашу внутреннюю функцию, которая делает всю магию
	// (У вас она называлась GenDiffRecursive)
	diff := gendiff.GenDiffRecursive(data1, data2)

	return diff, nil
}

// Эту функцию мы уже писали для main.go, ее можно перенести сюда,
// чтобы не дублировать код, а из main.go просто вызывать code.GenDiff(path1, path2)
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
