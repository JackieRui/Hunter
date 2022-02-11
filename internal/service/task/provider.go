package main

import (
	"git.hunter.net/hunter/config"
	"git.hunter.net/hunter/internal/dao/log"
	"git.hunter.net/hunter/internal/dao/task"
	"github.com/google/wire"
)

/*
TaskManager Provider
*/

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.New,
	task.ProviderSet,
)
