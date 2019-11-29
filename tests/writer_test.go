package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/helper"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Writer(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)

	for _, parent := range packages {
		for _, file := range parent.Node().Files {

			path, err := ioutil.TempFile(".", "prefix")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(path.Name())
			assert.Nil(helper.Save(file, path.Name()))
		}

	}
	//os.RemoveAll(path)

}
