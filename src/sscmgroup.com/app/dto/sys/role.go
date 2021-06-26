package sys

import (
	"sscmgroup.com/app/dto"
)

type RoleForm struct {
	RoleId  int    `form:"roleId" json:"roleId"`
	Title   string `form:"title" json:"title" binding:"required"`
	Name    string `form:"name" json:"name" binding:"required"`
	Remark  string `form:"remark" json:"remark"` // 备注
	Status  int    `form:"status" json:"status"` // 状态
	MenuIds []int  `form:"menuIds" json:"menuIds"`
}

type SysRoleSearch struct {
	dto.Pagination `search:"-"`
	RoleId         int    `form:"roleId"`
	RoleName       string `form:"roleName"`
	BeginTime      string `form:"beginTime"`
	EndTime        string `form:"endTime"`
	Status         string `form:"status"`
}
