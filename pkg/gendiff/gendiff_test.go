package gendiff

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiffJSON(t *testing.T) {
	// 1. Получаем путь к папке с тестами
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename)
	fixturesDir := filepath.Join(testDir, "testdata", "fixtures")

	// 2. Пути к файлам (строки)
	file1 := filepath.Join(fixturesDir, "file1.json")
	file2 := filepath.Join(fixturesDir, "file2.json")
	expectedFile := filepath.Join(fixturesDir, "stylish_result.txt")

	// 3. Вызываем GenDiff ПРЯМО с путями (парсинг внутри самой функции)
	// УБРАЛИ вызовы parser.ReadFile, так как GenDiff делает это сам
	actual, err := GenDiff(file1, file2, "stylish")
	require.NoError(t, err)

	// 4. Читаем ожидаемый результат для сравнения
	expected, err := os.ReadFile(expectedFile)
	require.NoError(t, err)

	// 5. Сравниваем
	assert.Equal(t, string(expected), actual)
}
