package database_standard

import "testing"

func TestSelectMyDataAll(t *testing.T) {
	SelectMyDataAll()
}

func TestSelectMyData(t *testing.T) {
	expect := MyData{1, "Taro", "taro@yamada", 39}

	md := SelectMyData(1)

	if md != expect {
		t.Errorf("not match instance. %v", md)
	}
}
