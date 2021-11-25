package task

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.hunter.net/hunter/internal/dao/job"
	"git.hunter.net/hunter/internal/utils"
)

/*
应届生招聘网任务Task
*/

type YJSTask struct {
	Task
}

func NewYJSTask(url, code, name string, retry int) *YJSTask {
	return &YJSTask{
		Task{
			Url:   url,
			Code:  code,
			Retry: retry,
			Name:  name,
		},
	}
}

// 获取网页详情
func (task *YJSTask) Run(ctx context.Context, ch chan<- ITask) (string, error) {
	// 构造请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", task.Url, nil)
	if err != nil {
		fmt.Printf("Task Request error:%v", err)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Add("referer", "https://www.yingjiesheng.com/zhuanye/jisuanji/hebei/")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Add("sec-ch-ua-platform", "macOS")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("request error:%v", err)
		return "", nil
	}
	defer resp.Body.Close()
	// 解析内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil readall error:%v", err)
		return "", nil
	}
	result := utils.ConvertToString(string(content), "GBK", "UTF-8")
	fmt.Printf("content:%s", result)
	// 判断分支
	return "", nil

}

func (task *YJSTask) Parse(ctx context.Context, response string) (*job.Job, error) {
	return nil, nil
}
