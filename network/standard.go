package network_standard

import (
	"fmt"
	"io"
	"net/http"
)

func Get(loc string) string {
	re, err := http.Get(loc)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	defer re.Body.Close()

	s, err := io.ReadAll(re.Body)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	return string(s)
}
