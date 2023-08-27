package file_standard

import (
	"os"
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

func TestReadFileALL(t *testing.T) {
	b, err := ReadFileAll("./testdata/sample.txt")
	if err != nil {
		t.Errorf("failed read file. error detail: %v", err)
	}

	text := string(b)

	if !strings.Contains(text, "sample file data") || !strings.Contains(text, "これはサンプルファイルです") {
		t.Errorf("not match text. file=%s", text)
	}
}

func TestReadFileLineByLine(t *testing.T) {
	b, err := ReadFileLineByLine("./testdata/sample.txt")
	if err != nil {
		t.Errorf("failed read file. error detail: %v", err)
	}

	text := string(b)

	if !strings.Contains(text, "sample file data") || !strings.Contains(text, "これはサンプルファイルです") {
		t.Errorf("not match text. file=%s", text)
	}
}

func TestWriteFile(t *testing.T) {
	d := t.TempDir()
	t.Logf("outdir: %s", d)

	str := "write this file by Golang!"
	data := []byte(str)

	err := WriteFile(d+"/"+"test.txt", data)
	if err != nil {
		t.Errorf("failed file create. error detail: %v", err)
	}

	f, err := os.Stat(d + "/" + "test.txt")
	if os.IsNotExist(err) {
		t.Errorf("don't exixts file. create failed. error detail=%v", err)
	}

	txt, _ := ReadFile(d + "/" + f.Name())
	if string(txt) != "write this file by Golang!" {
		t.Errorf("text not match! file=%s", string(txt))
	}
}
