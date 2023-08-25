package file_standard

import (
	"os"
)

func ReadFile(fp string) ([]byte, error) {
	return os.ReadFile(fp)
}
