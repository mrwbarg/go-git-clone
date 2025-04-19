package utils

import (
	"os"
	"path/filepath"
)

func FindRepoRoot(path string) (string, error) {
	gitdir := filepath.Join(path, ".git")

	fileInfo, err := os.Stat(gitdir)
	if err == nil {
		if fileInfo.IsDir() {
			return path, nil
		}
	}

	if path == "/" {
		return "", os.ErrNotExist
	}

	parentDir := filepath.Dir(path)
	if parentDir == path {
		return "", os.ErrNotExist
	}

	return FindRepoRoot(parentDir)

}
