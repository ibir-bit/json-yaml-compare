package main

import (
	"flag"
	"fmt"
	"os"

	"code/pkg/gendiff" // Убедитесь, что здесь ваш правильный импорт пакета

	"gopkg.in/yaml.v3"
)

func main() {
	// 1. Настраиваем флаг формата
	format := flag.String("format", "stylish", "set format of output")
	flag.Parse()

	// 2. Проверяем, что передано достаточно аргументов (два файла)
	if flag.NArg() < 2 {
		fmt.Println("Usage: gendiff [-format <format>] <firstConfig> <secondConfig>")
		os.Exit(1)
	}

	// 3. Получаем пути к файлам
	path1 := flag.Arg(0)
	path2 := flag.Arg(1)

	// Чтобы линтер не ругался на неиспользуемый format (пока мы не реализуем другие форматы)
	_ = format

	// 4. Читаем и парсим первый файл
	data1, err := readAndParse(path1)
	if err != nil {
		fmt.Printf("Error reading first file: %v\n", err)
		os.Exit(1)
	}

	// 5. Читаем и парсим второй файл
	data2, err := readAndParse(path2)
	if err != nil {
		fmt.Printf("Error reading second file: %v\n", err)
		os.Exit(1)
	}

	// 6. Генерируем diff (здесь только 1 возвращаемое значение!)
	diff := gendiff.GenDiffRecursive(data1, data2)

	// 7. Печатаем результат
	fmt.Println(diff)
}

// readAndParse — вспомогательная функция для чтения файла и парсинга YAML
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
