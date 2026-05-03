package internal

import "testing"

func TestFormat(t *testing.T) {
	diffs := []Diff2{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "-", V1: 2, Wrap: []Diff2{
			{Key: "a2", Type: " ", V1: 5},
			{Key: "b2", Type: " ", V1: 6, Wrap: []Diff2{
				{Key: "a4", Type: " ", V1: nil},
			}},
		}},
		{Key: "c", Type: "+", V2: 3},
		{Key: "d", Type: "+/-", V1: 4, V2: 40},
	}

	got := Format(diffs, 0)

	want := `{
    a: 1
  - b: {
        a2: 5
        b2: {
            a4: null
        }
    }
  + c: 3
  - d: 4
  + d: 40
}`

	if got != want {
		t.Errorf("unexpected:\n%s\n\ngot:\n%s", want, got)
	}
}
