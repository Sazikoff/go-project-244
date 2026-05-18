package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	
	cases := []struct {
		name   string
		file1  string
		file2  string
		format string

		want    string
		wantErr bool
	}{
		{
			name:   "json stylish",
			file1:  "testdata/fixture/file1.json",
			file2:  "testdata/fixture/file2.json",
			format: "stylish",
			want:   read(t, "testdata/fixture/expectedJson_stylish.txt"),
		},
		{
			name:   "yaml plain",
			file1:  "testdata/fixture/file3.yaml",
			file2:  "testdata/fixture/file4.yaml",
			format: "plain",
			want:   read(t, "testdata/fixture/expectedYaml_plain.txt"),
		},
		{
			name:    "file not exist",
			file1:   "testdata/fixture/not_exist.json",
			file2:   "testdata/fixture/file2.json",
			format:  "stylish",
			wantErr: true,
		},
	}

	for _, tc := range cases {

		t.Run(tc.name, func(t *testing.T) {
			got, err := GenDiff(tc.file1, tc.file2, tc.format)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func read(t *testing.T, path string) string {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read fixture: %s", path)
	}
	return string(b)
}
