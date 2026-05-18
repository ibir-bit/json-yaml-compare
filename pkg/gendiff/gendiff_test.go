package gendiff_test

import (
	"code/pkg/gendiff"
	"code/pkg/parser"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Вспомогательная функция для получения пути к фикстурам
func getFixturePath(filename string) string {
	return filepath.Join("testdata", "fixture", filename)
}

// Вспомогательная функция для чтения ожидаемого результата из файла
func readFixture(t *testing.T, filename string) string {
	t.Helper()
	content, err := os.ReadFile(getFixturePath(filename))
	require.NoError(t, err)
	return string(content)
}

// 1. Тест для формата STYLISH (читает из существующего файла-фикстуры)
func TestGenDiffNestedYAML(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	actual, err := gendiff.GenDiff(data1, data2, "stylish")
	require.NoError(t, err)

	expected := strings.TrimSpace(readFixture(t, "stylish_result.txt"))
	cleanActual := strings.TrimSpace(actual)

	assert.Equal(t, expected, cleanActual)
}

// 2. Тест для формата PLAIN (строка зашита в код)
func TestGenDiffPlain(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	actual, err := gendiff.GenDiff(data1, data2, "plain")
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

// 3. Тест для формата JSON (структурированная строка зашита в код через ` `)
func TestGenDiffJSON(t *testing.T) {
	file1 := getFixturePath("file1.yml")
	file2 := getFixturePath("file2.yml")

	data1, err := parser.ReadFile(file1)
	require.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	require.NoError(t, err)

	actual, err := gendiff.GenDiff(data1, data2, "json")
	require.NoError(t, err)

	expected := strings.TrimSpace(`[
    {
        "Key": "common",
        "Status": "nested",
        "Value": null,
        "OldValue": null,
        "NewValue": null,
        "Children": [
            {
                "Key": "follow",
                "Status": "added",
                "Value": false,
                "OldValue": null,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "setting1",
                "Status": "unchanged",
                "Value": "Value 1",
                "OldValue": null,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "setting2",
                "Status": "removed",
                "Value": 200,
                "OldValue": null,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "setting3",
                "Status": "changed",
                "Value": null,
                "OldValue": true,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "setting4",
                "Status": "added",
                "Value": "blah blah",
                "OldValue": null,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "setting5",
                "Status": "added",
                "Value": {
                    "key5": "value5"
                },
                "OldValue": null,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "setting6",
                "Status": "nested",
                "Value": null,
                "OldValue": null,
                "NewValue": null,
                "Children": [
                    {
                        "Key": "doge",
                        "Status": "nested",
                        "Value": null,
                        "OldValue": null,
                        "NewValue": null,
                        "Children": [
                            {
                                "Key": "wow",
                                "Status": "changed",
                                "Value": null,
                                "OldValue": "",
                                "NewValue": "so much",
                                "Children": null
                            }
                        ]
                    },
                    {
                        "Key": "key",
                        "Status": "unchanged",
                        "Value": "value",
                        "OldValue": null,
                        "NewValue": null,
                        "Children": null
                    },
                    {
                        "Key": "ops",
                        "Status": "added",
                        "Value": "vops",
                        "OldValue": null,
                        "NewValue": null,
                        "Children": null
                    }
                ]
            }
        ]
    },
    {
        "Key": "group1",
        "Status": "nested",
        "Value": null,
        "OldValue": null,
        "NewValue": null,
        "Children": [
            {
                "Key": "baz",
                "Status": "changed",
                "Value": null,
                "OldValue": "bas",
                "NewValue": "bars",
                "Children": null
            },
            {
                "Key": "foo",
                "Status": "unchanged",
                "Value": "bar",
                "OldValue": null,
                "NewValue": null,
                "Children": null
            },
            {
                "Key": "nest",
                "Status": "changed",
                "Value": null,
                "OldValue": {
                    "key": "value"
                },
                "NewValue": "str",
                "Children": null
            }
        ]
    },
    {
        "Key": "group2",
        "Status": "removed",
        "Value": {
            "abc": 12345,
            "deep": {
                "id": 45
            }
        },
        "OldValue": null,
        "NewValue": null,
        "Children": null
    },
    {
        "Key": "group3",
        "Status": "added",
        "Value": {
            "deep": {
                "id": {
                    "number": 45
                }
            },
            "fee": 100500
        },
        "OldValue": null,
        "NewValue": null,
        "Children": null
    }
]`)

	cleanActual := strings.TrimSpace(actual)
	assert.Equal(t, expected, cleanActual)
}
