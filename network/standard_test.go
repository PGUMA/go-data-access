package network_standard

import (
	"testing"
)

func TestGet(t *testing.T) {
	resb := Get("https://www.google.com/")
	if resb == "" {
		t.Errorf("can not get contents.")
	}
	t.Logf("contents: %s", resb)
}

func TestJsonGet(t *testing.T) {
	resj := JsonGet("https://tuyano-dummy-data.firebaseio.com/mydata.json")

	for _, c := range resj {
		s := c.Str()
		t.Logf("%v", s)
	}
}
