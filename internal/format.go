package internal

import (
	"fmt"
	"strings"
)

// Format converts a slice of differences into a human-readable string representation
func Format(diffs []Diff2, n int) string {
	var b strings.Builder
	var v1, v2 string
    
	b.WriteString("{\n")
    s := strings.Repeat("    ", n) // тут  4 пробела
    
	for _, d := range diffs {
        
		if d.Wrap != nil {
			v1 = Format(d.Wrap, n+1)
			v2 = Format(d.Wrap, n+1)
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
