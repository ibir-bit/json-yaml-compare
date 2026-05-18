package gendiff

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"code/pkg/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiffJSON(t *testing.T) {
	// 1. Получаем путь к папке, где лежит этот тест
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename)

	// 2. Указываем правильный путь к фикстурам
	fixturesDir := filepath.Join(testDir, "testdata", "fixtures")

	// 3. Файлы для сравнения
	file1 := filepath.Join(fixturesDir, "file1.json")
	file2 := filepath.Join(fixturesDir, "file2.json")
	expectedFile := filepath.Join(fixturesDir, "stylish_result.txt")

	// 4. Читаем данные через ваш парсер
	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	// 5. Генерируем результат (формат "stylish", так как ожидаемый файл называется stylish_result.txt)
	actual, err := GenDiff(data1, data2, "stylish")
	require.NoError(t, err)

	// 6. Читаем ожидаемый результат
	expected, err := os.ReadFile(expectedFile)
	require.NoError(t, err)

	// 7. Сравниваем (для текстового формата достаточно привести к строке)
	assert.Equal(t, string(expected), actual)
}
