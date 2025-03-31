package utils

import "os"

func IsDirectory(path string) (bool, error) {

	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}
