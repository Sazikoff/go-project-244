package formatters

import (
	"testing"
	."code/internal"
)

func TestFormatStylish(t *testing.T) {
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

	got := FormatStylish(diffs, 0)

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

func TestFormatPlain(t *testing.T) {
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

	got := FormatPlain(diffs)

	want :=
`Property 'b' was removed
Property 'c' was added with value: '3'
Property 'd' was updated. From '4' to '40'`

	if got != want {
		t.Errorf("unexpected:\n%s\n\ngot:\n%s", want, got)
	}
}
