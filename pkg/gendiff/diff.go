package gendiff

import "sort"

// DiffNode — элемент внутреннего представления diff
type DiffNode struct {
	Key      string
	Status   string      // added, removed, unchanged, changed, nested
	Value    interface{} // для added/removed/unchanged
	OldValue interface{} // для changed
	NewValue interface{} // для changed
	Children []DiffNode  // для nested
}

// buildDiff строит внутреннее дерево diff рекурсивно
func buildDiff(data1, data2 map[string]interface{}) []DiffNode {
	keysMap := make(map[string]struct{})
	for k := range data1 {
		keysMap[k] = struct{}{}
	}
	for k := range data2 {
		keysMap[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var diff []DiffNode
	for _, k := range keys {
		v1, ok1 := data1[k]
		v2, ok2 := data2[k]

		switch {
		case ok1 && ok2:
			m1, isMap1 := v1.(map[string]interface{})
			m2, isMap2 := v2.(map[string]interface{})

			if isMap1 && isMap2 {
				diff = append(diff, DiffNode{
					Key:      k,
					Status:   "nested",
					Children: buildDiff(m1, m2),
				})
			} else if v1 == v2 {
				diff = append(diff, DiffNode{
					Key:    k,
					Status: "unchanged",
					Value:  v1,
				})
			} else {
				diff = append(diff, DiffNode{
					Key:      k,
					Status:   "changed",
					OldValue: v1,
					NewValue: v2,
				})
			}
		case ok1 && !ok2:
			diff = append(diff, DiffNode{
				Key:    k,
				Status: "removed",
				Value:  v1,
			})
		case !ok1 && ok2:
			diff = append(diff, DiffNode{
				Key:    k,
				Status: "added",
				Value:  v2,
			})
		}
	}
	return diff
}
