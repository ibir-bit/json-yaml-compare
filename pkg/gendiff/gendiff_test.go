package gendiff_test

import (
	"code/pkg/gendiff"
	"testing"
)

func TestGenDiffRecursive(t *testing.T) {
	data1 := map[string]interface{}{"key": "value"}
	data2 := map[string]interface{}{"key": "new value"}

	// ВАЖНО: вызываем функцию через префикс 'gendiff.'
	result := gendiff.FormatStylish(data1, data2)

	if result == "" {
		t.Error("Expected result, got empty string")
	}
}
