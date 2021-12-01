package task

import (
	"context"

	"git.hunter.net/hunter/internal/dao/job"
)

/*
Task 每个页面抓取抽象成Task
*/

type ITask interface {
	// Task抓取列表页面资源 并判断是否继续抓取
	RunList(ctx context.Context, ch chan<- ITask)
	// Task抓取详情页资源 调用各自的解析方法对页面进行解析 返回最终的解析数据
	RunDetail(ctx context.Context, ch chan<- ITask) (*job.Job, error)
}

type Task struct {
	Url   string `json:"url"`   // 抓取链接
	Page  int    `json:"page"`  // 抓取第几页
	Type  int    `json:"type"`  // 类型 0:列表页 1:详情页 列表页数据未存储到库
	Retry int    `json:"retry"` // 重试次数
	Code  string `json:"code"`  // 抓取资源代码
	Name  string `json:"name"`  // 资源中文名称
}

type Config struct {
	Url   string `yaml:"url"`
	Code  string `yaml:"code"`
	Retry int    `yaml:"retry"`
	Name  string `yaml:"name"`
}
