package repository

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mrwbarg/go-git-clone/internal/models/config"
	"github.com/mrwbarg/go-git-clone/internal/models/object"
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

func (r *Repository) WriteObject(obj object.Object) error {
	hash := obj.Hash()

	path, err := r.file(true, "objects", hash[:2], hash[2:])
	if err != nil {
		return fmt.Errorf("fatal: error writing object %s: %v", hash, err)
	}

	var fileBuffer bytes.Buffer
	writer := zlib.NewWriter(&fileBuffer)
	_, err = writer.Write(obj.Serialize())
	if err != nil {
		return fmt.Errorf("fatal: error compressing object %s: %v", hash, err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("fatal: error compressing object %s: %v", hash, err)
	}

	err = os.WriteFile(path, fileBuffer.Bytes(), os.ModePerm)
	if err != nil {
		return fmt.Errorf("fatal: error writing object %s: %v", hash, err)
	}

	return nil
}

func (r *Repository) ReadObject(sha string) (*object.Object, error) {
	path, err := r.file(false, "objects", sha[:2], sha[2:])
	if err != nil {
		return nil, err
	}

	isFile, _ := utils.IsFile(path)
	if !isFile {
		return nil, fmt.Errorf("fatal: object %s not found", sha)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fatal: error reading object %s: %v", sha, err)
	}

	reader, err := zlib.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("fatal: error decompressing object %s: %v", sha, err)
	}
	err = file.Close()
	if err != nil {
		return nil, fmt.Errorf("fatal: error closing object %s: %v", sha, err)
	}

	uncompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("fatal: error reading object %s: %v", sha, err)
	}
	err = reader.Close()
	if err != nil {
		return nil, fmt.Errorf("fatal: error closing object %s: %v", sha, err)
	}

	obj, err := object.Load(uncompressed)
	if err != nil {
		return nil, fmt.Errorf("fatal: error creating object %s: %v", sha, err)
	}

	return obj, nil

}

func WithPath(path string, skipValidation bool) func(*Repository) {

	rootPath, err := utils.FindRepoRoot(path)
	if err != nil {
		if skipValidation {
			rootPath = path
		} else {
			utils.ErrorAndExit(fmt.Sprintf("fatal: not a git repository: %s", path))
		}
	}

	gitdir := filepath.Join(rootPath, ".git")

	return func(r *Repository) {
		r.worktree = path
		r.gitdir = gitdir

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

	_, err := repo.dir(true, "branches")
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error initializing the repo: %v", err))
	}

	_, err = repo.dir(true, "objects")
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error initializing the repo: %v", err))
	}

	_, err = repo.dir(true, "refs", "tags")
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error initializing the repo: %v", err))
	}

	_, err = repo.dir(true, "refs", "heads")
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error initializing the repo: %v", err))
	}

	err = os.WriteFile(
		repo.path("description"),
		[]byte("Unnamed repository; edit this file 'description' to name the repository.\n"),
		os.ModePerm,
	)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error initializing the repo: %v", err))
	}

	err = os.WriteFile(
		repo.path("HEAD"),
		[]byte("ref: refs/heads/master\n"),
		os.ModePerm,
	)
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error initializing the repo: %v", err))
	}

	repo.conf.Initialize(repo.gitdir)

	return repo
}

func (r *Repository) FindObject(name string, format *object.ObjectType, follow bool) string {
	return name
}

func (r *Repository) Log(start object.Commit) string {

	var builder strings.Builder
	builder.WriteString(start.Log())
	builder.WriteString("\n\n")

	parent := start.Parent()
	for parent != "" {
		obj, err := r.ReadObject(start.Parent())
		if err != nil {
			utils.ErrorAndExit(fmt.Sprintf("fatal: error reading object %s: %v", start.Parent(), err))
		}

		if (*obj).Type() != object.CommitType {
			utils.ErrorAndExit(fmt.Sprintf("fatal: object %s is not a commit", start.Parent()))
		}

		commit, ok := (*obj).(*object.Commit)
		if !ok {
			utils.ErrorAndExit(fmt.Sprintf("fatal: object %s is not a commit", start.Parent()))
		}

		builder.WriteString(commit.Log())
		builder.WriteString("\n\n")

		parent = commit.Parent()
	}

	return builder.String()

}
