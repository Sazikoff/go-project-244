package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildDiff(t *testing.T) {
	data1 := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"e": map[string]any{
			"e1": 4,
			"e2": map[string]any{
				"e3": nil,
			},
		},
		"g": map[string]any{
			"g1": 5,
			"g2": map[string]any{
				"g3": nil,
			},
		},
	}

	data2 := map[string]any{
		"a": 1,
		"b": 20,
		"d": 4,
		"f": map[string]any{}, //
		"g": map[string]any{
			"g1": 6,
			"g2": map[string]any{
				"g4": 7,
			},
		},
	}

	got := BuildDiff(data1, data2)

	want := []Diff2{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "+/-", V1: 2, V2: 20},
		{Key: "c", Type: "-", V1: 3},
		{Key: "d", Type: "+", V2: 4},
		{Key: "e", Type: "-", Wrap: []Diff2{
			{Key: "e1", Type: " ", V1: 4},
			{Key: "e2", Type: " ", Wrap: []Diff2{
				{Key: "e3", Type: " ", V1: nil},
			}},
		}},
		{Key: "f", Type: "+", Wrap: nil}, //
		{Key: "g", Type: " ", Wrap: []Diff2{
			{Key: "g1", Type: "+/-", V1: 5, V2: 6},
			{Key: "g2", Type: " ", Wrap: []Diff2{
				{Key: "g3", Type: "-", V1: nil},
				{Key: "g4", Type: "+", V2: 7},
			}},
		}},
	}

	assert.Equal(t, got, want)

}
