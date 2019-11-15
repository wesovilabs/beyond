package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/writer"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Writer(t *testing.T){
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)


	for _,parent:=range packages{
		for _,file:=range parent.Node().Files{

			path, err := ioutil.TempFile(".", "prefix")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(path.Name())
			assert.Nil(writer.Save(file,path.Name()))
		}

	}
	//os.RemoveAll(path)

}
