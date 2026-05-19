package gendiff_test

import (
	"code/pkg/gendiff" // Импортируем ваш пакет, чтобы тесты его "видели"
	"testing"
)

func TestGenDiffRecursive(t *testing.T) {
	// Подготовьте тестовые данные (ваши map[string]interface{})
	data1 := map[string]interface{}{"key": "value"}
	data2 := map[string]interface{}{"key": "new value"}

	// ВАЖНО: вызываем функцию через префикс 'gendiff.'
	result := gendiff.FormatStylish(data1, data2)

	// Пример проверки (используйте testify, если он у вас подключен)
	if result == "" {
		t.Error("Expected result, got empty string")
	}
}
