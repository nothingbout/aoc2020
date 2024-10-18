package lib

import (
	"testing"
	"testing/fstest"
)

func TestReadInputLines(t *testing.T) {
	data := `1721
979
366
299
675
1456`
	fs := fstest.MapFS{
		"files/example.txt": {Data: []byte(data)},
	}

	got := ReadInputLines(fs, "files/example.txt")
	want := []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
	AssertEqual(t, got, want)
}
