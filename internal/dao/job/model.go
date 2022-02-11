package job

import "time"

type Job struct {
	ID        int32     `json:"id" gorm:"primary_key"`               // 主键
	Title     string    `json:"title" gorm:"column:title"`           // 职位名称
	Company   string    `json:"company" gorm:"column:company"`       // 公司名称
	Content   string    `json:"content" gorm:"column:content"`       // 页面内容
	Url       string    `json:"url" gorm:"column:url"`               // 页面链接
	Status    int8      `json:"status" gorm:"column:status"`         // 职位状态 0:初始状态
	Task      string    `json:"task" gorm:"column:task"`             // Task信息
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (Job) TableName() string {
	return "job"
}
