package file_standard

import (
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	b, err := ReadFile("./testdata/sample.txt")
	if err != nil {
		t.Errorf("failed read file. error detail: %v", err)
	}

	text := string(b)

	if !strings.Contains(text, "sample file data") || !strings.Contains(text, "これはサンプルファイルです") {
		t.Errorf("not match text. file=%s", text)
	}
}