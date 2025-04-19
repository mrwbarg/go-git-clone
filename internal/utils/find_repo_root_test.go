package utils

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindRepoRoot_CurrentDir(t *testing.T) {
	testDir := t.TempDir()
	_ = os.MkdirAll(path.Join(testDir, ".git"), os.ModePerm)

	repoRoot, err := FindRepoRoot(testDir)
	assert.NoError(t, err)
	assert.Equal(t, testDir, repoRoot)
}

func Test_FindRepoRoot_OneDirUp(t *testing.T) {
	testDir := t.TempDir()
	_ = os.MkdirAll(path.Join(testDir, ".git"), os.ModePerm)

	repoRoot, err := FindRepoRoot(path.Join(testDir, "subdir"))
	assert.NoError(t, err)
	assert.Equal(t, testDir, repoRoot)
}

func Test_FindRepoRoot_TwoDirsUp(t *testing.T) {
	testDir := t.TempDir()
	_ = os.MkdirAll(path.Join(testDir, ".git"), os.ModePerm)

	repoRoot, err := FindRepoRoot(path.Join(testDir, "subdir", "another"))
	assert.NoError(t, err)
	assert.Equal(t, testDir, repoRoot)
}

func Test_FindRepoRoot_NoGitDir(t *testing.T) {
	testDir := t.TempDir()

	repoRoot, err := FindRepoRoot(testDir)
	assert.Error(t, err)
	assert.Equal(t, "", repoRoot)
}
