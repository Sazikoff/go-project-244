package internal

import "sort"

type Diff struct {
	Key  string
	Type string // "-", "+", " ", "+/-"
	V1   any
	V2   any
}

// ComputeDiff — временная заглушка.
// Просто возвращает первый аргумент без изменений.
func ComputeDiff(data1, data2 map[string]any) []Diff {

    // объединение мапов в слайс, сортировка
	sortKeys := unionKeys(data1, data2)

    // создание списка структур Diff
	result := make([]Diff, 0, len(sortKeys))

    // создание стркутр, заполнение ими списка
	for _, k := range sortKeys {
		v1, ok1 := data1[k]
		v2, ok2 := data2[k]

		switch {
		case ok1 && !ok2:
            result = append(result, Diff{Key: k, Type: "-", V1: v1})
			// слева есть, справа нет
		case !ok1 && ok2:
            result = append(result, Diff{Key: k, Type: "+", V2: v2})
			// справа есть, слева нет
		case ok1 && ok2:
			if v1 == v2 {
                result = append(result, Diff{Key: k, Type: " ", V1: v1})
				// без изменений
                } else {
                result = append(result, Diff{Key: k, Type: "+/-", V1: v1, V2: v2})
				// изменен
			}

		}
        
	}

    return result

}

func unionKeys(m1, m2 map[string]any) []string {
	keysMap := make(map[string]struct{})

	for k := range m1 {
		keysMap[k] = struct{}{}
	}
	for k := range m2 {
		keysMap[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
