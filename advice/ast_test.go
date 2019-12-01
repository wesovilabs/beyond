package advice

import "testing"

func Test_unsupportedType(t *testing.T) {
	if res := unsupportedType("sample"); res != "" {
		t.Fatalf("unexpected value")
	}
}
