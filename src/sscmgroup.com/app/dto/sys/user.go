package sys

import (
	"sscmgroup.com/app/dto"
)

type UserForm struct {
	UserId   int    `form:"userId" json:"userId"`
	UserName string `form:"userName" json:"userName"`
	RealName string `form:"realName" json:"realName"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	Avatar   string `form:"avatar" json:"avatar"`
	Status   int    `form:"status" json:"status"`
	RoleIds  []int  `form:"roleIds" json:"roleIds"`
}

type ChangePwdForm struct {
	OldPassword string `form:"oldPassword" json:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required"`
}

type LoginForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	//Code string `form:"code" json:"code"`
	UUID string `form:"uuid" json:"uuid" binding:"required"`
}

type AuthUser struct {
	Id       int    `json:"userId"`
	UserName string `json:"userName"`
	Roles    []string
}

type Auth struct {
	IsAuth bool
	User   AuthUser
}

type SysUserSearch struct {
	dto.Pagination `search:"-"`
	UserId         int    `form:"userId" json:"userId"`
	UserName       string `form:"userName" json:"userName"`
	Email          string `form:"email" json:"email"`
	Status         int    `form:"status" json:"status"`
}
