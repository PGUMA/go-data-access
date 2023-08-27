package network_standard

import "testing"

func TestGet(t *testing.T) {
	resb := Get("https://www.google.com/")
	if resb == "" {
		t.Errorf("can not get contents.")
	}
	t.Logf("contents: %s", resb)
}
