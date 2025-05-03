package commands

import (
	"fmt"
	"os"

	"github.com/mrwbarg/go-git-clone/internal/models/object"
	"github.com/mrwbarg/go-git-clone/internal/models/repository"
	"github.com/mrwbarg/go-git-clone/internal/utils"
)

func LogCmd(sha string) {
	dir, err := os.Getwd()
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: failed to get current working directory: %v", err))
	}
	repo := repository.New(repository.WithPath(dir, false))
	obj, err := repo.ReadObject(sha)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: %v", err))
	}

	commit, ok := (*obj).(*object.Commit)
	if !ok {
		utils.ErrorAndExit("fatal: not a commit object")
	}

	fmt.Println(repo.Log(*commit))
}
