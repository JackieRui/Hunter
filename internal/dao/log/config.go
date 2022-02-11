package log

import "go.uber.org/zap/zapcore"

/*
配置文件结构
*/

// Config serializes log
type Config struct {
	Level      string        `yaml:"level"`      // 日志级别
	LevelColor bool          `yaml:"levelColor"` // 日志级别字段开启颜色功能
	Format     string        `yaml:"format"`     // Log format
	Stdout     bool          `yaml:"stdout"`     // 是否输出到控制台
	File       FileLogConfig `yaml:"file"`       // File log config
}

type FileLogConfig struct {
	Filename   string `yaml:"filename"`   // 日志文件路径
	LogRotate  bool   `yaml:"logRotate"`  // log rotate enable
	MaxSize    int    `yaml:"maxSize"`    // Single File Max Size
	MaxDays    int    `yaml:"maxDays"`    // Max Log Keep Days
	MaxBackups int    `yaml:"maxBackups"` // Max Number of Old Log File
	BufSize    int    `yaml:"bufSize"`    // Max Size of bufio.Writer
}

// level 获取日志级别，默认是Info
func (c *Config) level() zapcore.Level {
	level := zapcore.InfoLevel
	if c.Level == "" {
		return level
	} else {
		if err := level.Set(c.Level); err != nil {
			panic(err)
		}
		return level
	}
}
