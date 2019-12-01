package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAvailableImportName(t *testing.T) {
	assert := assert.New(t)
	res := findAvailableImportName(map[string]string{
		"pkg/release": "release",
	}, "release")
	assert.Equal("_release", res)
}

func Test_addImportSpec(t *testing.T) {
	if res := addImportSpec(nil, "", ""); res != nil {
		t.Fatal("Expected empty")
	}
}

func Test_findImportName(t *testing.T) {
	res := findImportName(nil, "myname", "pkg/myname")
	if res != "myname" {
		t.Fatal("expcted myname")
	}
}
