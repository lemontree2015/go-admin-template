package sys

import (
	"time"
)

type Setting struct {
	Id      int       `gorm:"column:id;primary_key" json:"id"`
	Name    string    `gorm:"column:name" json:"name"`
	Title   string    `gorm:"column:title" json:"title"`
	Content string    `gorm:"column:content" json:"content"`
	Status  int       `gorm:"column:status" json:"status"`
	Created time.Time `gorm:"column:created" json:"created"`
	Updated time.Time `gorm:"column:updated" json:"updated"`
}

func (s *Setting) TableName() string {
	return "sys_settings"
}
