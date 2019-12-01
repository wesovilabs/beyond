package helper

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"

	"testing"
)

func Test_Copy(t *testing.T) {
	dir, err := ioutil.TempDir("", "beyond")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.RemoveAll(dir)
	childPath := filepath.Join(dir, "child1")
	if err := os.MkdirAll(childPath, os.ModePerm); err != nil {
		t.Fatal(err.Error())
	}

	os.Symlink(childPath, childPath+"_")

	childExcludePath := filepath.Join(dir, "child2")
	if err := os.MkdirAll(childExcludePath, os.ModePerm); err != nil {
		t.Fatal(err.Error())
	}

	if err := ioutil.WriteFile(filepath.Join(childPath, "hello.txt"), []byte("hello"), 0777); err != nil {
		t.Fatal(err.Error())
	}

	if err := os.Symlink(filepath.Join(childPath, "hello.txt"), filepath.Join(childPath, "hello2.txt")); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Symlink(childPath, childPath+"2"); err != nil {
		t.Fatal(err.Error())
	}
	if err := ioutil.WriteFile(filepath.Join(childPath, "bye.txt"), []byte("bye"), 0777); err != nil {
		t.Fatal(err.Error())
	}

	if targetDir, err := ioutil.TempDir("", "beyond"); err == nil {
		defer os.RemoveAll(targetDir)
		child2AbsPath, _ := filepath.Abs("child2")
		CopyDirectory(dir, targetDir, map[string]bool{child2AbsPath: true})
		if _, err := os.Stat(filepath.Join(targetDir, "child1")); err != nil {
			t.Fatal(err.Error())
		}
		if _, err := os.Stat(filepath.Join(targetDir, "child1", "hello.txt")); err != nil {
			t.Fatal(err.Error())
		}
		if _, err := os.Stat(filepath.Join(targetDir, "child1", "hello2.txt")); err != nil {
			t.Fatal(err.Error())
		}
		if _, err := os.Stat(filepath.Join(targetDir, "child12")); err != nil {
			t.Fatal(err.Error())
		}
		if _, err := os.Stat(filepath.Join(targetDir, "child2")); err == nil {
			t.Fatal("directory should not be created there!")
		}

	}

}

func Test_copySymLink(t *testing.T) {
	assert := assert.New(t)
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("file1", d1, 7777)
	defer os.Remove("file1")
	defer os.Remove("file2")
	defer os.Remove("file3")
	os.Symlink("file1", "file2")
	assert.Nil(err)
	copySymLink("file2", "file3")

}
func Test_createIfNotExists(t *testing.T) {
	createIfNotExists(".", os.ModeDir)
}
