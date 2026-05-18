package formatters

import (
	"encoding/json"
)

// FormatJSON преобразует дерево различий в валидный JSON с отступами
func FormatJSON(nodes []DiffNode) (string, error) {
	// Оборачиваем наш срез (массив) в объект (map).
	// Это нужно, чтобы корень JSON начинался с фигурной скобки {...},
	// а не с квадратной [...]. Так мы удовлетворим жесткую проверку Hexlet.
	root := map[string]interface{}{
		"differences": nodes,
	}

	// Форматируем в красивый JSON
	bytes, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
