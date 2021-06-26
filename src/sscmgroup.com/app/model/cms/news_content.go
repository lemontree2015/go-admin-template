package cms

import "time"

type NewsContent struct {
	Id      int       `gorm:"column:id;primary_key" json:"contentId"`
	NewsId  int       `gorm:"column:news_id" json:"newsId"`
	Content string    `gorm:"column:content" json:"content"`
	Status  int       `gorm:"column:status" json:"status"`
	Created time.Time `gorm:"column:created" json:"created"`
	Updated time.Time `gorm:"column:updated" json:"updated"`
}

func (m *NewsContent) TableName() string {
	return "news_contents"
}
