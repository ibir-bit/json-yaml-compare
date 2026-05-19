package gendiff // Тоже gendiff, чтобы находиться в одном пакете

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiff(t *testing.T) {
	type testCase struct {
		name         string
		file1        string
		file2        string
		format       string
		expectedFile string
	}

	tests := []testCase{
		{
			name:         "JSON to Stylish",
			file1:        "testdata/fixtures/file1.json",
			file2:        "testdata/fixtures/file2.json",
			format:       "stylish",
			expectedFile: "testdata/fixtures/stylish_result.txt",
		},
		{
			name:         "YAML to Stylish",
			file1:        "testdata/fixtures/file1.yml",
			file2:        "testdata/fixtures/file2.yml",
			format:       "stylish",
			expectedFile: "testdata/fixtures/stylish_result.txt",
		},
		{
			name:         "JSON to Plain",
			file1:        "testdata/fixtures/file1.json",
			file2:        "testdata/fixtures/file2.json",
			format:       "plain",
			expectedFile: "testdata/fixtures/plain_result.txt",
		},
		{
			name:         "YAML to Plain",
			file1:        "testdata/fixtures/file1.yml",
			file2:        "testdata/fixtures/file2.yml",
			format:       "plain",
			expectedFile: "testdata/fixtures/plain_result.txt",
		},
		{
			name:         "JSON to JSON",
			file1:        "testdata/fixtures/file1.json",
			file2:        "testdata/fixtures/file2.json",
			format:       "json",
			expectedFile: "testdata/fixtures/json_result.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedBytes, err := os.ReadFile(tt.expectedFile)
			require.NoError(t, err, "Failed to read expected file")
			expected := string(expectedBytes)

			// Вызываем функцию напрямую, без префиксов пакета
			actual, err := GenDiff(tt.file1, tt.file2, tt.format)

			require.NoError(t, err, "GenDiff returned an error")

			if tt.format == "json" {
				assert.JSONEq(t, expected, actual, "JSON structures do not match")
			} else {
				assert.Equal(t, expected, actual, "Diff output does not match expected")
			}
		})
	}
}
