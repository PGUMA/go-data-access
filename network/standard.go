package network_standard

import (
	"encoding/json"
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

type JsonResponse interface{}
type SampleResponse struct {
	Name string
	Mail string
	Tel  string
}

func (sr *SampleResponse) Str() string {
	return sr.Name + "|" + sr.Mail + "|" + sr.Tel
}

func JsonGet(loc string) []SampleResponse {
	re, err := http.Get(loc)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	defer re.Body.Close()

	s, err := io.ReadAll(re.Body)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	var data []SampleResponse
	err = json.Unmarshal(s, &data)
	if err != nil {
		fmt.Printf("json parse error. error detail=%v", err)
	}

	return data
}
