package helper

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

const goModFileName = "go.mod"

// GetModuleName returns the module name from a go.mod
func GetModuleName(rootDir string) (string, error) {
	goModPath := filepath.Join(rootDir, goModFileName)
	if _, err := os.Stat(goModPath); err != nil {
		return "", errors.New("module name cannot be found")
	}

	f, err := os.Open(goModPath)
	CheckError(err)

	defer closeFile(f)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(trimmedLine, "module") {
			return trimmedLine[7:], nil
		}
	}

	return "", errors.New("module name cannot be found")
}
