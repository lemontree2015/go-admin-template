package sys

import "sscmgroup.com/app/dto"

type SettingForm struct {
	Id      int    `form:"id" json:"id"`
	Name    string `form:"name" json:"name" binding:"required"`
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Status  int    `form:"status" json:"status"`
}

type SettingSearch struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"`
	BeginTime      string `form:"beginTime"`
	EndTime        string `form:"endTime"`
	Status         int    `form:"status"`
}
