package file_standard

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(fp string) ([]byte, error) {
	// 事前にFileサイズを検証してその通りにバッファサイズを確保してから読み込むのでReadAllより高速な模様
	return os.ReadFile(fp)
}

func ReadFileAll(fp string) ([]byte, error) {
	f, err := os.OpenFile(fp, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}

	defer f.Close()

	return io.ReadAll(f)
}

func ReadFileLineByLine(fp string) ([]byte, error) {
	f, err := os.Open(fp)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}

	defer f.Close()

	bu := bufio.NewReaderSize(f, 1024)

	var b = []byte{}
	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		}
		b = append(b, line...)
	}
	return b, nil
}

func WriteFile(fp string, b []byte) error {
	f, err := os.Create(fp)
	if err != nil {
		fmt.Printf("error %v", err)
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		fmt.Printf("error %v", err)
		return err
	}

	return nil
}
