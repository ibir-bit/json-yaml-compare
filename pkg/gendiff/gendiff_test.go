package gendiff_test

import (
	"io/ioutil"
	"testing"

	"code/pkg/gendiff"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
)

// readYAML читает YAML-файл и возвращает map[string]interface{}
func readYAML(filePath string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &data)
	return data, err
}

func TestGenDiffNestedYAML(t *testing.T) {
	file1 := "testdata/fixtures/file1.yml"
	file2 := "testdata/fixtures/file2.yml"

	data1, err := readYAML(file1)
	assert.NoError(t, err)

	data2, err := readYAML(file2)
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
