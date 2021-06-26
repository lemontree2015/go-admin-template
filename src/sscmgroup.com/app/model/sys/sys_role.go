package sys

import (
	"time"
)

type Role struct {
	Id      int       `gorm:"column:id;primary_key" json:"roleId"`
	Name    string    `gorm:"column:name" json:"name"`
	Title   string    `gorm:"column:title" json:"title"`
	Remark  string    `gorm:"column:remark" json:"remark"`
	Status  int       `gorm:"column:status" json:"status"`
	Created time.Time `gorm:"column:created" json:"created"`
	Updated time.Time `gorm:"column:updated" json:"updated"`
	MenuIds []int     `json:"menuIds" gorm:"-"`
}

func (r *Role) TableName() string {
	return "sys_role"
}

// ToRoleIDMap 转换为角色ID映射
func ToRoleIDMap(r []*Role) map[int]*Role {
	m := make(map[int]*Role)
	for _, item := range r {
		m[item.Id] = item
	}
	return m
}

// ToRoleIDMap 转换为角色ID映射
func ToRoleIDs(roles []Role) (roleIds []int) {
	for _, r := range roles {
		roleIds = append(roleIds, r.Id)
	}
	return
}
