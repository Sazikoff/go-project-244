package formatters

import (
	"fmt"
	"slices"
	"strings"
	"code/internal"
)

// FormatStylish converts a slice of differences into a human-readable string representation
func FormatStylish(diffs []internal.Diff2, n int) string {
	var b strings.Builder
	var v1, v2 string

	b.WriteString("{\n")
	s := strings.Repeat("    ", n) // тут  4 пробела

	for _, d := range diffs {

		if d.Wrap != nil {
			v1 = FormatStylish(d.Wrap, n+1)
			v2 = FormatStylish(d.Wrap, n+1)
		} else {
			v1 = fmt.Sprintf("%v", d.V1)
			if v1 == "<nil>" {
				v1 = "null"
			}
			v2 = fmt.Sprintf("%v", d.V2)
			if v2 == "<nil>" {
				v2 = "null"
			}
		}

		switch d.Type {

		case "-":
			fmt.Fprintf(&b, "%s  - %s: %s\n", s, d.Key, v1) //тут по 2 пробела перед +/-

		case "+":
			fmt.Fprintf(&b, "%s  + %s: %s\n", s, d.Key, v2) //тут по 2 пробела перед +/-

		case " ":
			fmt.Fprintf(&b, "%s    %s: %s\n", s, d.Key, v1) //тут по 2 пробела перед +/-

		case "+/-":
			fmt.Fprintf(&b, "%s  - %s: %s\n", s, d.Key, v1) //тут по 2 пробела перед +/-
			fmt.Fprintf(&b, "%s  + %s: %s\n", s, d.Key, v2) //тут по 2 пробела перед +/-
		}
	}
	fmt.Fprintf(&b, "%s}", s)
	// b.WriteString("}")
	return b.String()
}

var mySlice = []string{"null", "true", "false"}

func FormatPlain(diffs []internal.Diff2, acc ...string) string {
	var path string
	var b strings.Builder
	var v1, v2 string

	for _, d := range diffs {

		var st1, st2 string

		if len(acc) == 0 {
			path = d.Key
		} else {
			path = acc[0] + "." + d.Key
		}

		if d.Wrap != nil && d.Type == " " {
			path = FormatPlain(d.Wrap, path)
			fmt.Fprint(&b, path)
		} else {
			v1 = fmt.Sprintf("%v", d.V1)
			if v1 == "<nil>" {
				v1 = "null"
			}
			v2 = fmt.Sprintf("%v", d.V2)
			if v2 == "<nil>" {
				v2 = "null"
			}

			if !slices.Contains(mySlice, v1) {
				st1 = "'"
			}

			if !slices.Contains(mySlice, v2) {
				st2 = "'"
			}

			if d.Type == "-" {
				fmt.Fprintf(&b, "Property '%s' was removed\n", path)
			}

			if d.Type == "+" {
				if d.Wrap != nil {
					fmt.Fprintf(&b, "Property '%s' was added with value: [complex value]\n", path)
				} else {
					fmt.Fprintf(&b, "Property '%s' was added with value: %s%s%s\n", path, st2, v2, st2)
				}
			}

			if d.Type == "+/-" {
				switch {
				case d.Wrap != nil && v1 != "null":
					fmt.Fprintf(&b, "Property '%s' was updated. From %s%s%s to [complex value]\n", path, st1, v1, st1)
				case d.Wrap != nil && v2 != "null":
					fmt.Fprintf(&b, "Property '%s' was updated. From [complex value] to %s%s%s\n", path, st2, v2, st2)
				default:
					fmt.Fprintf(&b, "Property '%s' was updated. From %s%s%s to %s%s%s\n", path, st1, v1, st1, st2, v2, st2)
				}
			}
		}
	}

	result := b.String()
	if strings.HasSuffix(result, "\n") && len(acc)== 0 {
		result = result[:len(result)-1]
	}

	return result
}
