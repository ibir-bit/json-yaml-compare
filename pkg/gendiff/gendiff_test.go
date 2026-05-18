package gendiff_test

import (
	"testing"

	"code/pkg/gendiff"
	"code/pkg/parser"

	"github.com/stretchr/testify/assert"
)

func TestGenDiffJSON(t *testing.T) {
	file1 := "testdata/fixture/file1.json"
	file2 := "testdata/fixture/file2.json"

	data1, err := parser.ReadFile(file1)
	assert.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	assert.NoError(t, err)

	diff := gendiff.GenDiff(data1, data2)

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	assert.Equal(t, expected, diff)
}

func TestGenDiffYAML(t *testing.T) {
	file1 := "testdata/fixture/file1.yml"
	file2 := "testdata/fixture/file2.yml"

	data1, err := parser.ReadFile(file1)
	assert.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	assert.NoError(t, err)

	diff := gendiff.GenDiff(data1, data2)

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	assert.Equal(t, expected, diff)
}
