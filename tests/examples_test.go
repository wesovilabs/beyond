package tests

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const examplesRepository = "https://github.com/wesovilabs/beyond-examples.git"

func TestPublishedExamples(t *testing.T) {
	dir, err := ioutil.TempDir("", "beyond-examples")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	command := fmt.Sprintf("clone %s %s", examplesRepository, dir)
	cmd := exec.Command("git", strings.Split(command, " ")...)
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

}
