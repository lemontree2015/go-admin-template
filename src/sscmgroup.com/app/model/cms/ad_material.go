package cms

import "time"

type AdMaterials struct {
	Id        int       `gorm:"column:id;primary_key" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Desc      string    `gorm:"column:desc" json:"desc"`
	Thumbnail string    `gorm:"column:thumbnail" json:"thumbnail"`
	Link      string    `gorm:"column:link" json:"link"`
	Position  int       `gorm:"column:position" json:"position"`
	Status    int       `gorm:"column:status" json:"status"`
	Created   time.Time `gorm:"column:created" json:"created"`
	Updated   time.Time `gorm:"column:updated" json:"updated"`
}

func (m *AdMaterials) TableName() string {
	return "ad_materials"
}
