package task

import (
	"context"
)

/*
Task 每个页面抓取抽象成Task
*/

type ITask interface {
	// Run Task任务运行
	Run(ctx context.Context, ch chan<- ITask)
	// RunList Task抓取列表页面资源 并判断是否继续抓取
	RunList(ctx context.Context, ch chan<- ITask)
	// RunDetail Task抓取详情页资源 调用各自的解析方法对页面进行解析 返回最终的解析数据
	RunDetail(ctx context.Context, ch chan<- ITask)
}

type Task struct {
	Url   string `json:"url"`   // 抓取链接
	Retry int    `json:"retry"` // 重试次数
	Code  string `json:"code"`  // 抓取资源代码
	Name  string `json:"name"`  // 资源中文名称
}
