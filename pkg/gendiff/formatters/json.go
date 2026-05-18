package formatters

import (
	"encoding/json"
)

// FormatJSON преобразует дерево различий в валидный JSON с отступами
func FormatJSON(nodes []DiffNode) (string, error) {
	// json.MarshalIndent делает красивый JSON с переносами строк и отступами (4 пробела)
	bytes, err := json.MarshalIndent(nodes, "", "    ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
