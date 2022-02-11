package config

import (
	"testing"

	"git.hunter.net/hunter/internal/dao/log"
)

func TestLoadConfig(t *testing.T) {
	c := LoadConfig("./primary.yml")
	logger, _ := log.New(c.Logger)
	logger.Error("Test Error Logger")
	logger.Debug("Test Debug Logger")
	logger.Warn("Test Warn Logger")
	logger.Info("Test Info Logger")
}
