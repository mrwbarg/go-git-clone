package commands

import (
	"github.com/mrwbarg/go-git-clone/internal/models/repository"
)

func InitCmd(path string) {
	repository.Initialize(path)
}
