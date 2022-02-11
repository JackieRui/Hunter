package main

import (
	"git.hunter.net/hunter/config"
	"github.com/google/wire"
)

func InitTaskManager(path config.ConfPath) *TaskManager {
	panic(wire.Build(providerSet))
}
