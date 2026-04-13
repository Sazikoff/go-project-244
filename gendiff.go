package code

import (
	"code/internal"
	// "encoding/json"
	// "errors"
	// "go.yaml.in/yaml/v3"
)

// GenDiff reads and parses two files, computes the differences between them,
// and returns the result formatted according to the specified format.
func GenDiff(filepath1, filepath2, format string) (string, error) {

	data1, err := internal.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := internal.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	diffs := internal.ComputeDiff(data1, data2)

	out := internal.Format(diffs)

	return out, nil
}
