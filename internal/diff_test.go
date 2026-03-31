package internal

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestComputeDiff(t *testing.T) {
	data1 := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	data2 := map[string]any{
		"a": 1,
		"b": 20,
		"d": 4,
	}

	got := ComputeDiff(data1, data2)

	want := []Diff{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "+/-", V1: 2, V2: 20},
		{Key: "c", Type: "-", V1: 3},
		{Key: "d", Type: "+", V2: 4},
	}

	assert.Equal(t, got, want)

}