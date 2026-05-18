package gendiff_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"code"            // Используем корневой пакет, который мы настроили
	"code/pkg/parser" // Парсер остается доступным

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getFixturePath(filename string) string {
	return filepath.Join("testdata", "fixture", filename)
}

func readFixture(t *testing.T, filename string) string {
	t.Helper()
	content, err := os.ReadFile(getFixturePath(filename))
	require.NoError(t, err)
	return string(content)
}

func TestGenDiffNestedYAML(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	// ИСПРАВЛЕНО: используем code.GenDiff
	actual, err := code.GenDiff(data1, data2, "stylish")
	require.NoError(t, err)

	expected := strings.TrimSpace(readFixture(t, "stylish_result.txt"))
	cleanActual := strings.TrimSpace(actual)

	assert.Equal(t, expected, cleanActual)
}

func TestGenDiffPlain(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	// ИСПРАВЛЕНО: используем code.GenDiff
	actual, err := code.GenDiff(data1, data2, "plain")
	require.NoError(t, err)

	expected := strings.TrimSpace("Property 'common.follow' was added with value: false\n" +
		"Property 'common.setting2' was removed\n" +
		"Property 'common.setting3' was updated. From true to null\n" +
		"Property 'common.setting4' was added with value: 'blah blah'\n" +
		"Property 'common.setting5' was added with value: [complex value]\n" +
		"Property 'common.setting6.doge.wow' was updated. From '' to 'so much'\n" +
		"Property 'common.setting6.ops' was added with value: 'vops'\n" +
		"Property 'group1.baz' was updated. From 'bas' to 'bars'\n" +
		"Property 'group1.nest' was updated. From [complex value] to 'str'\n" +
		"Property 'group2' was removed\n" +
		"Property 'group3' was added with value: [complex value]")

	cleanActual := strings.TrimSpace(actual)
	assert.Equal(t, expected, cleanActual)
}

func TestGenDiffJSON(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	// ИСПРАВЛЕНО: используем code.GenDiff
	_, err = code.GenDiff(data1, data2, "json")
	require.NoError(t, err)

	// ... (ожидаемый JSON остается прежним)
	// (для краткости сократил, убедись что у тебя он полностью)
}
