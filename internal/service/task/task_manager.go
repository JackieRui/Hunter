package main

import (
	"context"
	"time"

	"git.hunter.net/hunter/internal/dao/log"

	"git.hunter.net/hunter/internal/dao/task"
	"go.uber.org/zap"
)

type TaskManager struct {
	Logger  *zap.Logger   // 日志
	YJSTask *task.YJSTask // 抓取任务
}

func NewTaskManager(logger *zap.Logger, yjsTask *task.YJSTask) *TaskManager {
	return &TaskManager{
		Logger:  logger,
		YJSTask: yjsTask,
	}
}

// InitProducer 初始化生产者 将抓取Task放入通道内
func (tm *TaskManager) InitProducer(ch chan<- task.ITask) {
	tasks := []task.ITask{
		tm.YJSTask,
	}
	// 如果len(tasks)>len(ch) 该goroutine就会阻塞
	for i := range tasks {
		ch <- tasks[i]
	}
}

// Consumer 消费者消费Task
func (tm *TaskManager) Consumer(ctx context.Context, ch chan task.ITask) {
	for {
		select {
		// 从通道接收Task来进行处理
		case t := <-ch:
			go t.Run(ctx, ch)
		// 1分钟内通道无Task 退出循环
		case <-time.After(60 * time.Second):
			break
		}
	}
}

func (tm *TaskManager) Start() {
	// 上下文 用于传递数据
	ctx := context.Background()
	log.L(ctx).Info("TaskManager Start....")

	// 通道大小为10
	ch := make(chan task.ITask, 10)
	// 生产者
	go tm.InitProducer(ch)
	// 消费者
	go tm.Consumer(ctx, ch)

	log.L(ctx).Info("TaskManager Done")
}
