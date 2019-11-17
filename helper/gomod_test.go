package helper

import "testing"

func Test_GoMod(t *testing.T) {
	if module, err := GetModuleName("../"); err != nil {
		t.Fatal(err.Error())
	} else {
		if module != "github.com/wesovilabs/goa" {
			t.Fatal("unexpected module name")
		}
	}
}
