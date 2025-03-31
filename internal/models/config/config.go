package config

import (
	"fmt"
	"reflect"

	"github.com/mrwbarg/go-git-clone/internal/utils"
	"github.com/spf13/viper"
)

type coreConfig struct {
	RepositoryFormatVersion int  `mapstructure:"repositoryformatversion"`
	FileMode                bool `mapstructure:"filemode"`
	Bare                    bool `mapstructure:"bare"`
}

type Config struct {
	Core coreConfig `mapstructure:"core"`
}

func WithRepositoryFormatVersion(version int) func(*Config) {
	return func(c *Config) {
		c.Core.RepositoryFormatVersion = version
	}
}

func WithFileMode(fileMode bool) func(*Config) {
	return func(c *Config) {
		c.Core.FileMode = fileMode
	}
}

func WithBare(bare bool) func(*Config) {
	return func(c *Config) {
		c.Core.Bare = bare
	}
}

func New(options ...func(*Config)) *Config {
	c := &Config{
		Core: coreConfig{
			RepositoryFormatVersion: 0,
			FileMode:                false,
			Bare:                    false,
		},
	}

	for _, option := range options {
		option(c)
	}
	return c
}

func setDefaults(v *viper.Viper, config interface{}) {
	val := reflect.ValueOf(config)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		panic("config must be a pointer to a struct")
	}

	setDefaultsRecursive(v, val.Elem(), "")
}

func setDefaultsRecursive(v *viper.Viper, val reflect.Value, parentKey string) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Name

		var key string
		if parentKey != "" {
			key = parentKey + "." + fieldName
		} else {
			key = fieldName
		}

		if field.Kind() == reflect.Struct {
			setDefaultsRecursive(v, field, key)
		} else {
			v.SetDefault(key, field.Interface())
		}
	}
}

func (c *Config) Initialize(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	setDefaults(viper.GetViper(), c)

	err := viper.SafeWriteConfig()
	if err != nil {
		utils.ErrorAndExit(fmt.Sprintf("fatal: error writing configuration file: %v", err))
	}
}
