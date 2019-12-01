package tests

import (
	"fmt"
	"github.com/wesovilabs/beyond/internal"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"testing"
)

const examplesRepository = "https://github.com/wesovilabs/beyond-examples.git"

func TestPublishedExamples(t *testing.T) {
	dir, err := ioutil.TempDir("", "beyond-examples")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.RemoveAll(dir)
	cmd := exec.Command("git", "clone", examplesRepository,".")
	cmd.Dir=dir
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		t.Fatal(err)
	}


	settings:=&internal.Settings{
		Work:true,
		Verbose:true,
		OutputDir:filepath.Join(dir,"generated"),
		Path: filepath.Join(dir,"before"),
		Project: "github.com/wesovilabs/beyond-examples/before",
		ExcludeDirs:map[string]bool{
			"generated":true,
			".git":true,
		},
		Pkg:"cmd",
	}
	goCmd:=internal.GoCommand(settings,[]string{"run","cmd/main.go"}).Do()
	internal.ExecuteMain(goCmd,settings)

}
