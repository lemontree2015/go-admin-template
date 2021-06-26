package cms

import "time"

type NewsCategory struct {
	Id       int            `gorm:"column:id;primary_key" json:"cateId"`
	ParentId int            `gorm:"column:parent_id" json:"parentId"`
	Name     string         `gorm:"column:name" json:"name"`
	Status   int            `gorm:"column:status" json:"status"`
	Created  time.Time      `gorm:"column:created" json:"created"`
	Updated  time.Time      `gorm:"column:updated" json:"updated"`
	News     []News         `json:"news" gorm:"-"`
	Children []NewsCategory `json:"children" gorm:"-"`
}

func (m *NewsCategory) TableName() string {
	return "news_category"
}
