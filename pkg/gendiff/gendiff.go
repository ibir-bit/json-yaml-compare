package gendiff

// GenDiffRecursive собирает всё воедино
func GenDiffRecursive(data1, data2 map[string]interface{}) string {
	diffTree := BuildDiff(data1, data2)
	return FormatStylish(diffTree, 1)
}
