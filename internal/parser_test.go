package internal

import (
	"os"
	"testing"
)

func writeTempFile(t *testing.T, content, ext string) string {
	t.Helper()

	file, err := os.CreateTemp("", "test-*"+ext)
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		t.Fatal(err)
	}

	file.Close()

	return file.Name()
}

func TestParseFile_JSON(t *testing.T) {
	content := `{"a": 1, "b": "text"}`

	path := writeTempFile(t, content, ".json")
	defer os.Remove(path)

	res, err := ParseFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res["a"] != float64(1) {
		t.Errorf("expected 1, got %v", res["a"])
	}
}

func TestParseFile_YAML(t *testing.T) {
	content := `
a: 1
b: text
`

	path := writeTempFile(t, content, ".yml")
	defer os.Remove(path)

	res, err := ParseFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res["b"] != "text" {
		t.Errorf("expected text, got %v", res["b"])
	}
}

func TestParseFile_UnsupportedFormat(t *testing.T) {
	content := `hello world`

	path := writeTempFile(t, content, ".txt")
	defer os.Remove(path)

	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}