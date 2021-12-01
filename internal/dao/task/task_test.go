package task

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func justTestYJSTaskRun(t *testing.T) {
	ctx := context.Background()
	ch := make(chan ITask, 100)
	task := &YJSTask{
		Task{
			Url:   "https://www.yingjiesheng.com/job-005-548-358.html",
			Page:  0,
			Type:  1,
			Code:  "YJS",
			Retry: 3,
			Name:  "应届生求职网",
		},
	}
	fmt.Printf("task:%v", task)
	task.RunDetail(ctx, ch)
}

func TestParseYJSList(t *testing.T) {
	// file, _ := os.Open("./yjs_list_01.html")
	// content, _ := ioutil.ReadAll(file)
	// // tablePattern := regexp.MustCompile("<table id=\"tb_job_list\" class=\"jobul\">(?s:(.*?))</table>")
	// // 条目解析
	// // jobliPattern := regexp.MustCompile("<tr class=\"jobli\">(?s:(.*?))</tr>")
	// // joblis := jobliPattern.FindAllString(string(content), -1)
	// // 内容解析 解析子url
	// trPattern := regexp.MustCompile("<td width=\"329\">(?s:.*?)<a href=\"(?s:(.*?))\"(?s:.*?)<td width=\"92\"><span class=\"sub\">(?s:(.*?))</span>")
	// tds := trPattern.FindAllStringSubmatch(string(content), 1)
	// fmt.Println("tds:", len(tds))
	// fmt.Println(tds[0][1])
	// fmt.Println(tds[0][2])
	// fmt.Println("--------------------------")
	// for _, td := range tds {
	// 	fmt.Printf("1:%s\n", td[1])
	// 	fmt.Printf("2:%s\n", td[2])
	// 	// fmt.Printf("3:%s\n", td[3])
	// 	fmt.Println("--------------------------")
	// }
	file, _ := os.Open("./yjs_detail_01.html")
	content, _ := ioutil.ReadAll(file)
	divComtitPattern := regexp.MustCompile("<div class=\"comtit clear\">(?s:.*?)<h1>(?s:(.*?))</h1>(?s:(.*?))<div class=\"sp_msg\">(?s:.*?)" +
		"<div id=\"wordDiv\" class=\"reprintJob tborder\">(?s:(.*?))<ul class=\"linkbtn\">")
	divs := divComtitPattern.FindAllStringSubmatch(string(content), 1)
	fmt.Printf("h1:%s\n", divs[0][1])
	fmt.Printf("p:%s\n", divs[0][2])
	fmt.Printf("wordDiv:%v\n", divs[0][3])
}
