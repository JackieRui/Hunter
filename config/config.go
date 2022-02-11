package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"git.hunter.net/hunter/internal/dao/log"

	"git.hunter.net/hunter/internal/dao/task"
	"gopkg.in/yaml.v2"
)

type ConfPath string

type Config struct {
	Logger    log.Config  `yaml:"logger"`
	YJSConfig task.Config `yaml:"yjs_task"`
}

func LoadConfig(path ConfPath) Config {
	bytes := read(string(path))
	c := Config{}

	err := yaml.UnmarshalStrict(bytes, &c)
	if err != nil {
		panic(fmt.Errorf("Config Content Error %v", err))
	}
	return c
}

func read(path string) []byte {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(fmt.Errorf("Config Path Error %v", err))
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("Config File Error %v", err))
	}
	return bytes
}

func LoadYJSTaskConfig(conf Config) task.Config {
	return conf.YJSConfig
}

func LoadLoggerConfig(conf Config) log.Config {
	return conf.Logger
}
