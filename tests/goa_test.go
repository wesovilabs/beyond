package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/internal"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Goa(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	path, err := ioutil.TempDir(".", "prefix")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path)
	internal.Run(pkg, packages, path)
}
