package task

import "github.com/google/wire"

/*
配置要抓取的内容
*/

var ProviderSet = wire.NewSet(
	NewYJSTask,
)
