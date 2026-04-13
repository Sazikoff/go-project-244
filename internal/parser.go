package internal

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"go.yaml.in/yaml/v3"
)

// ParseFile reads and parses a JSON or YAML file into a map.
func ParseFile(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		// json.Unmarshal
		err = json.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}
	case ".yaml", ".yml":
		// yaml.Unmarshal
		err = yaml.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported format")
	}

	return result, nil
}