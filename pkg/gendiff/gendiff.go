package gendiff

// GenDiffRecursive — публичная функция для рекурсивного diff
func GenDiffRecursive(data1, data2 map[string]interface{}) string {
	diffTree := buildDiff(data1, data2)
	return formatStylish(diffTree, 1)
}
