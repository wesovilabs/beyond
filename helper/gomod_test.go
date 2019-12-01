package helper

import "testing"

func Test_GoMod(t *testing.T) {
	if module, err := GetModuleName("../"); err != nil {
		t.Fatal(err.Error())
	} else {
		if module != "github.com/wesovilabs/beyond" {
			t.Fatal("unexpected module name")
		}
	}
	if module, err := GetModuleName("."); err == nil {
		t.Fatalf("unexpected module name")
	} else if module != "" {
		t.Fatalf("unexpected module name")
	}
}
