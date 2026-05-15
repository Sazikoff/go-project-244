package code

import (
	"code/formatters"
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

	diffs := internal.BuildDiff(data1, data2)

	var out string

	switch format {
	case "stylish":
		out = formatters.FormatStylish(diffs, 0)
	case "plain":
		out = formatters.FormatPlain(diffs)
	}



	return out, nil
}
