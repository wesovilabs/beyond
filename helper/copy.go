package helper

import (
	"fmt"
	"github.com/wesovilabs/goa/logger"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

// CopyDirectory function that copies directories
func CopyDirectory(scrDir, dest string, excludes map[string]string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}

	for index := range entries {
		entry := entries[index]

		if _, ok := excludes[entry.Name()]; ok {
			continue
		}

		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		if err:=copy(fileInfo, sourcePath, destPath, excludes);err!=nil{
			return err
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}

	return nil
}

func copy(fileInfo os.FileInfo, sourcePath, destPath string, excludes map[string]string) error {
	switch fileInfo.Mode() & os.ModeType {
	case os.ModeDir:
		if err := CreateIfNotExists(destPath, 0755); err != nil {
			return err
		}

		if err := CopyDirectory(sourcePath, destPath, excludes); err != nil {
			return err
		}
	case os.ModeSymlink:
		if err := CopySymLink(sourcePath, destPath); err != nil {
			return err
		}
	default:
		if err := Copy(sourcePath, destPath); err != nil {
			return err
		}
	}

	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	defer func(){
		if err:=out.Close();err!=nil{
			logger.Error(err.Error())
		}
	}()

	if err != nil {
		return err
	}

	in, err := os.Open(srcFile)

	defer func() {
		if closeErr := in.Close(); closeErr != nil {
			logger.Errorf(closeErr.Error())
		}
	}()

	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}

	return os.Symlink(link, dest)
}