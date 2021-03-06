package helper

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CopyDirectory function that copies directories
func CopyDirectory(scrDir, dest string, excludes map[string]bool) {
	entries, err := ioutil.ReadDir(scrDir)
	CheckError(err)

	for index := range entries {
		entry := entries[index]
		copyEntry(scrDir, dest, excludes, entry)
	}
}

func copyEntry(scrDir, dest string, excludes map[string]bool, entry os.FileInfo) {
	entryAbsPath, _ := filepath.Abs(entry.Name())

	if _, ok := excludes[entryAbsPath]; ok {
		return
	}

	sourcePath := filepath.Join(scrDir, entry.Name())
	destPath := filepath.Join(dest, entry.Name())

	if exists(destPath) {
		return
	}

	fileInfo, err := os.Stat(sourcePath)
	CheckError(err)

	copy(fileInfo, sourcePath, destPath, excludes)

	isSymlink := entry.Mode()&os.ModeSymlink != 0
	if !isSymlink {
		err := os.Chmod(destPath, entry.Mode())
		CheckError(err)
	}
}

func copy(fileInfo os.FileInfo, sourcePath, destPath string, excludes map[string]bool) {
	switch fileInfo.Mode() & os.ModeType {
	case os.ModeDir:
		createIfNotExists(destPath, 0755)
		CopyDirectory(sourcePath, destPath, excludes)
	case os.ModeSymlink:
		copySymLink(sourcePath, destPath)
	default:
		Copy(sourcePath, destPath)
	}
}

// Copy copy directory yo other
func Copy(srcFile, dstFile string) {
	out, err := os.Create(dstFile)

	defer closeFile(out)

	CheckError(err)

	in, err := os.Open(srcFile)

	defer closeFile(in)
	CheckError(err)

	_, err = io.Copy(out, in)
	CheckError(err)
}

func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func createIfNotExists(dir string, perm os.FileMode) {
	if exists(dir) {
		return
	}

	err := os.MkdirAll(dir, perm)
	CheckError(err)
}

func copySymLink(source, dest string) {
	link, err := os.Readlink(source)
	CheckError(err)
	CheckError(os.Symlink(link, dest))
}
