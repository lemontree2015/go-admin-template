package cms

import "time"

type News struct {
	Id        int       `gorm:"column:id;primary_key" json:"id"`
	CateId    int       `gorm:"column:cate_id" json:"cateId"`
	Title     string    `gorm:"column:title" json:"title"`
	Desc      string    `gorm:"column:desc" json:"desc"`
	Thumbnail string    `gorm:"column:thumbnail" json:"thumbnail"`
	NewsType  int       `gorm:"column:news_type" json:"newsType"`
	Status    int       `gorm:"column:status" json:"status"`
	Content   string    `json:"content" gorm:"-"`
	Created   time.Time `gorm:"column:created" json:"created"`
	Updated   time.Time `gorm:"column:updated" json:"updated"`
}

func (m *News) TableName() string {
	return "news_titles"
}
