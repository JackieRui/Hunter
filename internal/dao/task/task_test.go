package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

// func TestYJSTaskRun(t *testing.T) {
// 	// ctx := context.Background()
// 	// ch := make(chan ITask, 100)
// 	var (
// 		url   = "https://www.yingjiesheng.com/hebeijob/list_1.html"
// 		code  = "YJS"
// 		retry = 3
// 		name  = "应届生求职网"
// 	)
// 	task := NewYJSTask(url, code, name, retry)
// 	fmt.Printf("task:%v", task)
// 	// task.RunList(ctx, ch)
// }

func TestParseYJSList(t *testing.T) {
	file, _ := os.Open("./yjs_list_01.html")
	content, _ := ioutil.ReadAll(file)
	// tablePattern := regexp.MustCompile("<table id=\"tb_job_list\" class=\"jobul\">(?s:(.*?))</table>")
	// 条目解析
	// jobliPattern := regexp.MustCompile("<tr class=\"jobli\">(?s:(.*?))</tr>")
	// joblis := jobliPattern.FindAllString(string(content), -1)
	// 内容解析 解析子url
	trPattern := regexp.MustCompile("<tr class=\"jobli\">(?s:.*?)<td width=\"329\">(?s:.*?)<a href=\"(?s:(.*?))\" target=\"_blank\">(?s:(.*?))</a>(?s:.*?)")
	tds := trPattern.FindAllStringSubmatch(string(content), -1)
	fmt.Println("tds:", len(tds))
	for _, td := range tds {
		fmt.Printf("1:%s\n", td[1])
		fmt.Printf("2:%s\n", td[2])
		// fmt.Printf("3:%s\n", td[3])
		fmt.Println("--------------------------")
	}
}
