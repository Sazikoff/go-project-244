package code

import (
	"code/internal/diff"
	"code/internal/parser"
	"encoding/json"
	"errors"
	"go.yaml.in/yaml/v3"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {

	data1, err := parser.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parser.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	diff := diff.ComputeDiff(data1, data2)

	var out []byte

	switch format {
	case "json":
		out, err = json.MarshalIndent(diff, "", "  ")

	case "yaml":
		out, err = yaml.Marshal(diff)
	default:
		return "", errors.New("unsupported output format")
	}

	if err != nil {
		return "", err
	}

	out, err = json.MarshalIndent(diff, "", "  ")

	return string(out), nil
}
