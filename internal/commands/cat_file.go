package commands

import (
	"fmt"
	"os"

	"github.com/mrwbarg/go-git-clone/internal/models/repository"
	"github.com/mrwbarg/go-git-clone/internal/utils"
)

func CatFileCmd(sha string) {
	dir, err := os.Getwd()
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: failed to get current working directory: %v", err))
	}
	repo := repository.New(repository.WithPath(dir, false))
	object, err := repo.ReadObject(sha)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: %v", err))
	}

	fmt.Print(string((*object).Content()))
}
