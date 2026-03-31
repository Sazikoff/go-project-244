package code

import (
	"os"
	"testing"
)

func TestGenDiff(t *testing.T) {
	expected, err := os.ReadFile("testdata/fixture/expected.txt")
	if err != nil {
		t.Fatalf("failed to read expected file: %v", err)
	}

	cases := []struct {
		name    string
		file1   string
		file2   string
		format  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid json files",
			file1:   "testdata/fixture/file1.json",
			file2:   "testdata/fixture/file2.json",
			format:  "json",
			want:    string(expected),
			wantErr: false,
		},
		{
			name:    "file not exist",
			file1:   "testdata/fixture/not_exist.json",
			file2:   "testdata/fixture/file2.json",
			format:  "json",
			wantErr: true,
		},
		{
			name:    "invalid json",
			file1:   "testdata/fixture/bad.json",
			file2:   "testdata/fixture/file2.json",
			format:  "json",
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GenDiff(tc.file1, tc.file2, tc.format)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.want {
				t.Fatalf("diff mismatch\ngot:\n%s\nwant:\n%s", got, tc.want)
			}
		})
	}
}
