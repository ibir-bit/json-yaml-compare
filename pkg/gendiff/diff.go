package gendiff

import (
	"code/pkg/gendiff/formatters"
	"sort"
)

func BuildDiff(data1, data2 map[string]interface{}) []formatters.DiffNode {
	keys := getSortedKeys(data1, data2)
	var nodes []formatters.DiffNode

	for _, key := range keys {
		val1, ok1 := data1[key]
		val2, ok2 := data2[key]

		if ok1 && ok2 {
			map1, isMap1 := val1.(map[string]interface{})
			map2, isMap2 := val2.(map[string]interface{})

			if isMap1 && isMap2 {
				nodes = append(nodes, formatters.DiffNode{
					Key:      key,
					Status:   "nested",
					Children: BuildDiff(map1, map2),
				})
			} else if val1 == val2 {
				nodes = append(nodes, formatters.DiffNode{
					Key:    key,
					Status: "unchanged",
					Value:  val1,
				})
			} else {
				nodes = append(nodes, formatters.DiffNode{
					Key:      key,
					Status:   "changed",
					OldValue: val1,
					NewValue: val2,
				})
			}
		} else if ok1 {
			nodes = append(nodes, formatters.DiffNode{Key: key, Status: "removed", Value: val1})
		} else {
			nodes = append(nodes, formatters.DiffNode{Key: key, Status: "added", Value: val2})
		}
	}
	return nodes
}

func getSortedKeys(m1, m2 map[string]interface{}) []string {
	keysMap := make(map[string]bool)
	for k := range m1 {
		keysMap[k] = true
	}
	for k := range m2 {
		keysMap[k] = true
	}
	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
