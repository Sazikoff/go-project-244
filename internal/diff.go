package internal

import (
	"sort"
)

type Diff2 struct {
	Key  string
	Type string // "-", "+", " ", "+/-"
	V1   any
	V2   any
	Wrap []Diff2
}

// BuildDiff returns a sorted slice of unique keys from both maps m1 and m2
func BuildDiff(m1, m2 map[string]any) []Diff2 {

	// здеь создаем сам сет
	keysMap := make(map[string]struct{})

	for k := range m1 {
		keysMap[k] = struct{}{}
	}
	for k := range m2 {
		keysMap[k] = struct{}{}
	}
	// создаем список пустых структур, грубо говоря сетов
	// для того чтоб можно было это дело отсортировать
	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}

	// сортировка ключей
	sort.Strings(keys)
	var sliceDiff2 []Diff2
	for _, k := range keys {
		sliceDiff2 = append(sliceDiff2, HelperDiff(k, m1, m2))
	}

	// fmt.Println(keys)
	return sliceDiff2
}

// HelperDiff принимает ключ, проверяет наличие в обоих map-ах
// и в зависимости от этого выдает соответствующую карточку Diff2
func HelperDiff(k string, m1, m2 map[string]any) Diff2 {

	var result Diff2

	// проверка наличия в map-ах
	v1, ok1 := m1[k]
	v2, ok2 := m2[k]

	switch {
	case m1 == nil || m2 == nil:
		var v any
		if v1 == nil {
			v = v2
		} else {
			v = v1
		}
		if m, ok := v.(map[string]any); ok {
			result = Diff2{Key: k, Type: " ", Wrap: BuildDiff(m, nil)}
		} else {
			result = Diff2{Key: k, Type: " ", V1: v}
		}
	case ok1 && !ok2:
		if m1, ok := v1.(map[string]any); ok {
			result = Diff2{Key: k, Type: "-", Wrap: BuildDiff(m1, nil)}
		} else {
			result = Diff2{Key: k, Type: "-", V1: v1}
		}
	case !ok1 && ok2:
		if m2, ok := v2.(map[string]any); ok {
			result = Diff2{Key: k, Type: "+", Wrap: BuildDiff(nil, m2)}
		} else {
			result = Diff2{Key: k, Type: "+", V2: v2}
		}
	case ok1 && ok2:
		if m1, ok := v1.(map[string]any); ok {
			if m2, ok := v2.(map[string]any); ok {
				result = Diff2{Key: k, Type: " ", Wrap: BuildDiff(m1, m2)}
				return result
			}
			result = Diff2{Key: k, Type: "+/-", V2: v2, Wrap: BuildDiff(m1, nil)}

		} else if m2, ok := v2.(map[string]any); ok {
			result = Diff2{Key: k, Type: "+/-", V1: v1, Wrap: BuildDiff(nil, m2)}
		} else {
			if v1 == v2 {
				result = Diff2{Key: k, Type: " ", V1: v1}
			} else {
				result = Diff2{Key: k, Type: "+/-", V1: v1, V2: v2}
			}
		}
	}
	return result
}
