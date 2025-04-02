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

func WithPath(path string, skipValidation bool) func(*Repository) {
	gitdir := filepath.Join(path, "potato")
	isDir, _ := utils.IsDirectory(gitdir)

	if !skipValidation && !isDir {
		utils.ErrorAndExit(fmt.Sprintf("fatal: not a git repository: %s", path))
	}

	return func(r *Repository) {
		configPath, err := r.file(false, "config.toml")
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: error checking for config file: %v", err))
		}

		configPathExists, _ := utils.PathExists(configPath)
		if configPath != "" && configPathExists {
			r.conf.Load(filepath.Dir(configPath))
		} else if !skipValidation {
			utils.ErrorAndExit("fatal: no config file found")
		}

		if !skipValidation && r.conf.Core.RepositoryFormatVersion != 0 {
			utils.ErrorAndExit("fatal: unsupported repositoryformatversion")
		}

		r.worktree = path
		r.gitdir = gitdir
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

func Initialize(path string) *Repository {
	repo := New(WithPath(path, true))

	pathExists, _ := utils.PathExists(repo.worktree)
	if pathExists {
		isDir, _ := utils.IsDirectory(repo.worktree)
		if !isDir {
			utils.ErrorAndExit(fmt.Sprintf("fatal: %s is not a directory", repo.worktree))
		}

		isRepo, _ := utils.IsDirectory(repo.gitdir)
		if isRepo {
			utils.ErrorAndExit("fatal: repository already exists")
		}
	} else {
		err := os.MkdirAll(repo.worktree, os.ModePerm)
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: error creating directory %s: %v", repo.worktree, err))
		}
	}

	repo.dir(true, "branches")
	repo.dir(true, "objects")
	repo.dir(true, "refs", "tags")
	repo.dir(true, "refs", "heads")

	return repo
}
