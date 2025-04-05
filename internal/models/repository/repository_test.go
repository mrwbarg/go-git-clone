package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Repository_Path(t *testing.T) {
	repo := &Repository{
		gitdir: ".test-git-dir",
	}

	objectsPath := repo.path("objects")
	refsPath := repo.path("refs")

	assert.Equal(t, ".test-git-dir/objects", objectsPath)
	assert.Equal(t, ".test-git-dir/refs", refsPath)
}

func Test_Repository_Dir_DoesNotExist(t *testing.T) {
	tempDir := t.TempDir()
	repo := &Repository{
		gitdir: tempDir,
	}

	path, err := repo.dir(false, "objects")

	assert.Equal(t, "", path)
	assert.NoError(t, err)

	path, err = repo.dir(true, "objects")

	assert.Equal(t, tempDir+"/objects", path)
	assert.NoError(t, err)
}

func Test_Repository_Dir_Exists(t *testing.T) {
	tempDir := t.TempDir()
	repo := &Repository{
		gitdir: tempDir,
	}

	_ = os.Mkdir(tempDir+"/objects", os.ModePerm)

	path, err := repo.dir(false, "objects")
	assert.Equal(t, tempDir+"/objects", path)
	assert.NoError(t, err)

	path, err = repo.dir(true, "objects")
	assert.Equal(t, tempDir+"/objects", path)
	assert.NoError(t, err)
}

func Test_Repository_File_PathDoesNotExist(t *testing.T) {
	tempDir := t.TempDir()
	repo := &Repository{
		gitdir: tempDir,
	}

	path, err := repo.file(false, "objects", "HEAD")
	assert.Equal(t, "", path)
	assert.NoError(t, err)

	path, err = repo.file(true, "objects", "HEAD")
	assert.Equal(t, tempDir+"/objects/HEAD", path)
	assert.NoError(t, err)

	headFileInfo, err := os.Stat(tempDir + "/objects/HEAD")
	assert.Nil(t, headFileInfo)
	assert.Error(t, err)
}

func Test_Repository_File_PathExists(t *testing.T) {
	tempDir := t.TempDir()
	repo := &Repository{
		gitdir: tempDir,
	}

	_ = os.Mkdir(tempDir+"/objects", os.ModePerm)

	path, err := repo.file(false, "objects", "HEAD")
	assert.Equal(t, tempDir+"/objects/HEAD", path)
	assert.NoError(t, err)

	path, err = repo.file(true, "objects", "HEAD")
	assert.Equal(t, tempDir+"/objects/HEAD", path)
	assert.NoError(t, err)

	headFileInfo, err := os.Stat(tempDir + "/objects/HEAD")
	assert.Nil(t, headFileInfo)
	assert.Error(t, err)
}

func Test_Repository_Initialize(t *testing.T) {
	tempDir := t.TempDir()
	repo := Initialize(tempDir)

	assert.NotNil(t, repo)

	objectsPath, err := os.Stat(tempDir + "/.git/objects")
	assert.NoError(t, err)
	assert.True(t, objectsPath.IsDir())

	refsPath, err := os.Stat(tempDir + "/.git/refs")
	assert.NoError(t, err)
	assert.True(t, refsPath.IsDir())

	tagsPath, err := os.Stat(tempDir + "/.git/refs/tags")
	assert.NoError(t, err)
	assert.True(t, tagsPath.IsDir())

	headsPath, err := os.Stat(tempDir + "/.git/refs/heads")
	assert.NoError(t, err)
	assert.True(t, headsPath.IsDir())

	descriptionFile, err := os.Stat(tempDir + "/.git/description")
	assert.NoError(t, err)
	assert.False(t, descriptionFile.IsDir())

	headFile, err := os.Stat(tempDir + "/.git/HEAD")
	assert.NoError(t, err)
	assert.False(t, headFile.IsDir())
}
