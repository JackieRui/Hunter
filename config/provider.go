package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	LoadConfig,
	LoadLoggerConfig,
	LoadYJSTaskConfig,
)