package internal

import (
	"fmt"
	"strings"
)

// Format converts a slice of differences into a human-readable string representation
func Format(diffs []Diff) string {
    var b strings.Builder

    b.WriteString("{\n")

    for _, d := range diffs {
        switch d.Type {

        case "-":
            fmt.Fprintf(&b, "  - %s: %v\n", d.Key, d.V1)

        case "+":
            fmt.Fprintf(&b, "  + %s: %v\n", d.Key, d.V2)

        case " ":
            fmt.Fprintf(&b, "    %s: %v\n", d.Key, d.V1)

        case "+/-":
            fmt.Fprintf(&b, "  - %s: %v\n", d.Key, d.V1)
            fmt.Fprintf(&b, "  + %s: %v\n", d.Key, d.V2)
        }
    }

    b.WriteString("}")
    return b.String()
}