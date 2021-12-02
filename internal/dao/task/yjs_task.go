package task

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"git.hunter.net/hunter/internal/dao/job"
	"git.hunter.net/hunter/internal/utils"
)

/*
应届生招聘网任务Task
*/

const YJSPrefix = "https://www.yingjiesheng.com"

type YJSTask struct {
	Task
}

func NewYJSTask(url, code, name string, retry int) *YJSTask {
	return &YJSTask{
		Task{
			Url:   url,
			Page:  1,
			Type:  0,
			Code:  code,
			Retry: retry,
			Name:  name,
		},
	}
}

// 获取列表页数据
func (task *YJSTask) RunList(ctx context.Context, ch chan<- ITask) {
	if task.Retry <= 0 {
		// 记录日志 不再抓取此任务
		// return nil, utils.ErrorWithCodeMsg(utils.StatusRetryLimit, "重试次数已达上限")
		return
	}
	// 构造请求
	client := &http.Client{}
	var url = fmt.Sprintf(task.Url, task.Page)
	req, err := http.NewRequest("GET", url, nil)
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
	// 请求资源
	resp, err := client.Do(req)
	if err != nil {
		// 记录日志 重新抓取
		task.Retry -= 1
		return
	}
	defer resp.Body.Close()
	// 读取内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 记录日志 重新抓取
		task.Retry -= 1
		return
	}
	// 编码转换
	result := utils.ConvertToString(string(content), "GBK", "UTF-8")
	err = ioutil.WriteFile("./list1.html", []byte(result), 0666)
	if err != nil {
		fmt.Printf("写文件错误:%v", err)
	}

	// 解析内容 列表页解析
	jobliPattern := regexp.MustCompile("<tr class=\"jobli\">(?s:(.*?))</tr>")
	joblis := jobliPattern.FindAllString(string(content), -1)
	trPattern := regexp.MustCompile("<td width=\"329\">(?s:.*?)<a href=\"(?s:(.*?))\"(?s:.*?)<td width=\"92\"><span class=\"sub\">(?s:(.*?))</span>")
	for _, jobli := range joblis {
		// 只匹配一次
		tds := trPattern.FindAllStringSubmatch(jobli, 1)
		if len(tds) > 0 {
			currentDate := tds[0][2]
			// 只抓取当天的数据
			if currentDate == utils.CurrentDate() {
				var url = tds[0][1]
				if strings.Contains(url, "http") {
					url = YJSPrefix + url
				}
				// 构造详情页抓取task 发送到ch TODO
				detailTask := &YJSTask{
					Task{
						Url:   url,
						Page:  0,
						Type:  1,
						Code:  task.Code,
						Retry: 3,
						Name:  task.Name,
					},
				}
				fmt.Printf("%v", detailTask)
			} else {
				// 历史记录 直接返回 不再抓取解析
				return
			}
		}
	}
	// 构造下一列表页的抓取任务
	nextTask := &YJSTask{
		Task{
			Url:   fmt.Sprintf(task.Url, task.Page+1),
			Page:  task.Page + 1,
			Type:  task.Type,
			Code:  task.Code,
			Retry: 3,
			Name:  task.Name,
		},
	}
	fmt.Printf("%v", nextTask)
}

// 获取详情页数据
func (task *YJSTask) RunDetail(ctx context.Context, ch chan<- ITask) (*job.Job, error) {
	if task.Retry <= 0 {
		// 记录日志 不再抓取此任务
		return nil, utils.ErrorWithCodeMsg(utils.StatusRetryLimit, "重试次数已达上限")
	}
	// 构造请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", task.Url, nil)
	if err != nil {
		fmt.Printf("Task Request error:%v", err)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Add("referer", "https://www.yingjiesheng.com/hebeijob/index.html")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	req.Header.Add("sec-ch-ua-platform", "macOS")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	// 请求资源
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, nil
	}
	defer resp.Body.Close()
	// 读取内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, nil
	}
	// 编码转换
	result := utils.ConvertToString(string(content), "GBK", "UTF-8")
	divPattern := regexp.MustCompile("<div class=\"comtit clear\">(?s:.*?)<h1>(?s:(.*?))</h1>(?s:(.*?))<div class=\"sp_msg\">(?s:.*?)" +
		"<div id=\"wordDiv\" class=\"reprintJob tborder\">(?s:(.*?))<ul class=\"linkbtn\">")
	divs := divPattern.FindAllStringSubmatch(result, 1)
	if len(divs) > 0 {
		t, _ := json.Marshal(task)
		j := &job.Job{
			Company: divs[0][1],
			Title:   divs[0][2],
			Content: divs[0][3],
			Url:     task.Url,
			Status:  0,
			Task:    string(t),
		}
		return j, nil
	}
	return nil, nil
}
