package code

import (
	"code/internal"
	// "encoding/json"
	// "errors"
	// "go.yaml.in/yaml/v3"
)

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
