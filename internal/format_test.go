package internal

import "testing"

func TestFormat(t *testing.T) {
	diffs := []Diff{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "-", V1: 2},
		{Key: "c", Type: "+", V2: 3},
		{Key: "d", Type: "+/-", V1: 4, V2: 40},
	}

	got := Format(diffs)

	want := `{
    a: 1
  - b: 2
  + c: 3
  - d: 4
  + d: 40
}`

	if got != want {
		t.Errorf("unexpected:\n%s\n\ngot:\n%s", want, got)
	}
}