package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_New_WithNoOptions(t *testing.T) {
	config := New()

	assert.Equal(t, 0, config.Core.RepositoryFormatVersion)
	assert.Equal(t, false, config.Core.FileMode)
	assert.Equal(t, false, config.Core.Bare)
}

func Test_New_WithOptions(t *testing.T) {
	config := New(WithRepositoryFormatVersion(1), WithFileMode(true), WithBare(true))

	assert.Equal(t, 1, config.Core.RepositoryFormatVersion)
	assert.Equal(t, true, config.Core.FileMode)
	assert.Equal(t, true, config.Core.Bare)
}

func Test_SetDefaults(t *testing.T) {
	type LevelTwo struct {
		Arg21 string `mapstructure:"arg21"`
		Arg22 string `mapstructure:"arg22"`
	}
	type LevelOne struct {
		Arg11  string   `mapstructure:"arg11"`
		Level2 LevelTwo `mapstructure:"level2"`
	}

	v := viper.GetViper()

	setDefaults(v, &LevelOne{})
	assert.ElementsMatch(t, []string{"arg11", "level2.arg21", "level2.arg22"}, v.AllKeys())
}

func Test_Config_Initialize(t *testing.T) {
	config := New()
	tempDir := t.TempDir()
	config.Initialize(tempDir)

	fileInfo, err := os.Stat(tempDir + "/config.toml")
	assert.Equal(t, "config.toml", fileInfo.Name())
	assert.NoError(t, err)
}

func Test_Config_Load(t *testing.T) {
	tempDir := t.TempDir()

	existingConfig := New(
		WithRepositoryFormatVersion(1),
		WithFileMode(true),
		WithBare(true),
	)
	existingConfig.Initialize(tempDir)

	config := New()
	config.Load(tempDir)

	assert.Equal(t, 1, config.Core.RepositoryFormatVersion)
	assert.Equal(t, true, config.Core.FileMode)
	assert.Equal(t, true, config.Core.Bare)
}
