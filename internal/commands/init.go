package commands

import (
	"fmt"

	"github.com/mrwbarg/go-git-clone/internal/models/config"
)

func InitCmd() {
	fmt.Println("Initializing a new git repository")
	config := config.New()
	config.Initialize("./")
}
