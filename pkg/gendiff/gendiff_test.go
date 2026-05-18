package gendiff_test

import (
	"testing"

	"code/pkg/gendiff"
	"code/pkg/parser"

	"github.com/stretchr/testify/assert"
)

func TestGenDiffNestedYAML(t *testing.T) {
	file1 := "testdata/fixture/file1.yml"
	file2 := "testdata/fixture/file2.yml"

	data1, err := parser.ReadFile(file1)
	assert.NoError(t, err)

	data2, err := parser.ReadFile(file2)
	assert.NoError(t, err)

	diff := gendiff.GenDiffRecursive(data1, data2)

	expected := `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: <nil>
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow:
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`

	assert.Equal(t, expected, diff)
}
