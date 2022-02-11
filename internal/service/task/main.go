package main

import (
	"flag"

	"git.hunter.net/hunter/config"
)

/*
数据抓取入口
*/

func main() {
	cfgPath := flag.String("config", "", "配置文件路径")
	flag.Parse()

	taskManager := InitTaskManager(config.ConfPath(*cfgPath))
	taskManager.Start()
}
