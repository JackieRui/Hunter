package task

/*
任务配置
*/

type Config struct {
	Url   string `yaml:"url"`   // 抓取URL
	Code  string `yaml:"code"`  // 自定义CODE
	Retry int    `yaml:"retry"` // 重试次数
	Name  string `yaml:"name"`  // 资源中文名称
}
