package commands

import (
	"fmt"

	"github.com/mrwbarg/go-git-clone/internal/models/repository"
)

func InitCmd() {
	fmt.Println("Initializing a new git repository")
	repository.Initialize("./")
}
