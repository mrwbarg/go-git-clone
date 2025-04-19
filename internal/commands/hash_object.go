package commands

import (
	"fmt"
	"os"

	"github.com/mrwbarg/go-git-clone/internal/models/object"
	"github.com/mrwbarg/go-git-clone/internal/models/repository"
	"github.com/mrwbarg/go-git-clone/internal/utils"
)

func HashObjectCmd(objectType object.ObjectType, path string, write bool) {

	data, err := os.ReadFile(path)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: failed to read file: %v", err))
	}

	obj, err := object.New(objectType, data)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: failed to create object: %v", err))
	}

	if write {
		dir, err := os.Getwd()
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: failed to get current working directory: %v", err))
		}
		repo := repository.New(repository.WithPath(dir, false))
		err = repo.WriteObject(*obj)
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: error writing object: %v", err))
		}

	}
	fmt.Print((*obj).Hash())
}
