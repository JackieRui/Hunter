package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

// func TestYJSTaskRun(t *testing.T) {
// 	ctx := context.Background()
// 	ch := make(chan ITask, 100)
// 	var (
// 		url   = "https://www.yingjiesheng.com/hebeijob/list_1.html"
// 		code  = "YJS"
// 		retry = 3
// 		name  = "应届生求职网"
// 	)
// 	task := NewYJSTask(url, code, name, retry)
// 	task.RunList(ctx, ch)
// }

func TestParseYJSList(t *testing.T) {
	file, _ := os.Open("./yjs_list_01.html")
	content, _ := ioutil.ReadAll(file)
	// tablePattern := regexp.MustCompile("<table id=\"tb_job_list\" class=\"jobul\">(?s:(.*?))</table>")
	jobliPattern := regexp.MustCompile("<tr class=\"jobli\">(?s:(.*?))</tr>")
	joblis := jobliPattern.FindAllString(string(content), -1)
	for _, jobli := range joblis {
		fmt.Printf("%v\n", jobli)
		fmt.Println("----------------")
	}
}
