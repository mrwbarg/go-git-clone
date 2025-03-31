package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mrwbarg/go-git-clone/internal/utils"
)

type Repository struct {
	worktree string
	gitdir   string
	conf     string
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

	if !(pathExists) {
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

func (r *Repository) file(path ...string) (string, error) {
	dirPath, err := r.dir(false, path[:len(path)-1]...)
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

	// TODO: parse config
	conf := "not implemented"
	return func(r *Repository) {
		r.worktree = path
		r.gitdir = gitdir
		r.conf = conf
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
