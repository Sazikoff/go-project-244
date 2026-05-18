package formatters

import (
	"code/internal"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFormatStylish(t *testing.T) {
	diffs := []internal.Diff2{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "-", V1: 2, Wrap: []internal.Diff2{
			{Key: "a2", Type: " ", V1: 5},
			{Key: "b2", Type: " ", V1: 6, Wrap: []internal.Diff2{
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
	diffs := []internal.Diff2{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "-", V1: 2, Wrap: []internal.Diff2{
			{Key: "a2", Type: " ", V1: 5},
			{Key: "b2", Type: " ", V1: 6, Wrap: []internal.Diff2{
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

func TestFormatJson(t *testing.T) {
	diffs := []internal.Diff2{
		{Key: "a", Type: " ", V1: 1},
		{Key: "b", Type: "-", V1: 2, Wrap: []internal.Diff2{
			{Key: "a2", Type: " ", V1: 5},
			{Key: "b2", Type: " ", V1: 6, Wrap: []internal.Diff2{
				{Key: "a4", Type: " ", V1: nil},
			}},
		}},
		{Key: "c", Type: "+", V2: 3},
		{Key: "d", Type: "+/-", V1: 4, V2: 40},
	}

	got := FormatJson(diffs)

	want :=
		`[
  {
    "key": "a",
    "type": " ",
    "oldValue": 1,
    "newValue": null
  },
  {
    "key": "b",
    "type": "-",
    "oldValue": 2,
    "newValue": null,
    "children": [
      {
        "key": "a2",
        "type": " ",
        "oldValue": 5,
        "newValue": null
      },
      {
        "key": "b2",
        "type": " ",
        "oldValue": 6,
        "newValue": null,
        "children": [
          {
            "key": "a4",
            "type": " ",
            "oldValue": null,
            "newValue": null
          }
        ]
      }
    ]
  },
  {
    "key": "c",
    "type": "+",
    "oldValue": null,
    "newValue": 3
  },
  {
    "key": "d",
    "type": "+/-",
    "oldValue": 4,
    "newValue": 40
  }
]`

	assert.JSONEq(t, want, got)

}
