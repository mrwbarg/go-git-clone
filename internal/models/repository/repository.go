package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mrwbarg/go-git-clone/internal/models/config"
	"github.com/mrwbarg/go-git-clone/internal/utils"
)

type Repository struct {
	worktree string
	gitdir   string
	conf     config.Config
}

func (r *Repository) path(path ...string) string {
	args := []string{r.gitdir}
	args = append(args, path...)
	return filepath.Join(args...)
}

func (r *Repository) dir(make bool, path ...string) (string, error) {

	fullPath := r.path(path...)
	pathExists, err := utils.PathExists(fullPath)
	if err != nil {
		return "", err
	}

	if pathExists {
		isDir, err := utils.IsDirectory(fullPath)
		if err != nil {
			return "", err
		}

		if isDir {
			return fullPath, nil
		}

		return "", fmt.Errorf("fatal: not a directory: %s", fullPath)
	}

	if make {
		err := os.MkdirAll(fullPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		return fullPath, nil
	}

	return "", nil
}

func (r *Repository) file(make bool, path ...string) (string, error) {
	dirPath, err := r.dir(make, path[:len(path)-1]...)
	if err != nil {
		return "", err
	}

	if dirPath != "" {
		return r.path(path...), nil
	}

	return "", nil
}

func WithPath(path string, force bool) func(*Repository) {
	gitdir := filepath.Join(path, ".git")

	isDir, err := utils.IsDirectory(gitdir)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error checking if %s is a directory: %v", gitdir, err))
	}
	if !(force || isDir) {
		utils.ErrorAndExit(fmt.Sprintf("fatal: not a git repository: %s", path))
	}

	return func(r *Repository) {
		r.worktree = path
		r.gitdir = gitdir

		configPath, err := r.file(false, "config.toml")
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: error checking for config file: %v", err))
		}

		configPathExists, err := utils.PathExists(configPath)
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: error checking if %s exists: %v", configPath, err))
		}
		if configPath != "" && configPathExists {
			r.conf.Load(filepath.Dir(configPath))
		} else if !force {
			utils.ErrorAndExit("fatal: no config file found")
		}

		if !force && r.conf.Core.RepositoryFormatVersion != 0 {
			utils.ErrorAndExit("fatal: unsupported repositoryformatversion")
		}
	}
}

func New(options ...func(*Repository)) *Repository {
	repo := &Repository{}

	for _, option := range options {
		option(repo)
	}

	if repo.worktree == "" {
		utils.ErrorAndExit("fatal: you need to specify a path for the repository")
	}

	return repo
}

func (r *Repository) Create(path string) {

}
