package task

import (
	"context"
	"testing"
)

func TestYJSTaskRun(t *testing.T) {
	ctx := context.Background()
	ch := make(chan ITask, 100)
	var (
		url   = "https://www.yingjiesheng.com/hebeijob/list_1.html"
		code  = "YJS"
		retry = 3
		name  = "应届生求职网"
	)
	task := NewYJSTask(url, code, name, retry)
	task.Run(ctx, ch)
}
