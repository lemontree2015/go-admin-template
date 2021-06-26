package sys

import (
	"time"
)

type User struct {
	Id            int       `gorm:"column:id;primary_key" json:"userId"`
	UserName      string    `gorm:"column:user_name" json:"userName"`
	RealName      string    `gorm:"column:real_name" json:"realName"`
	Password      string    `gorm:"column:password" json:"-"`
	Email         string    `gorm:"column:email" json:"email"`
	Phone         string    `gorm:"column:phone" json:"phone"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	Status        int       `gorm:"column:status" json:"status"`
	LastLoginTime time.Time `gorm:"column:last_login_time;default:NULL" json:"lastLoginTime"`
	LastLoginIp   string    `gorm:"column:last_login_ip" json:"lastLoginIp"`
	Created       time.Time `gorm:"column:created" json:"created"`
	Updated       time.Time `gorm:"column:updated" json:"updated"`
	Roles         []Role    `json:"roles" gorm:"-"`
	Permissions   []string  `json:"permissions" gorm:"-"`
}

func (u *User) TableName() string {
	return "sys_user"
}
