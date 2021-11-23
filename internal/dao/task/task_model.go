package task

import (
	"context"

	"git.hunter.net/hunter/internal/dao/job"
)

/*
Task 每个任务抽象成Task
*/

type Task interface {
	// Task抓取页面资源 返回内容 并判断是否继续抓取
	Run(ctx context.Context, ch chan<- Task) (string, error)
	// 不同Task调用各自的解析方法对页面进行解析 返回最终的解析数据
	Parse(ctx context.Context, response string) (*job.Job, error)
}
