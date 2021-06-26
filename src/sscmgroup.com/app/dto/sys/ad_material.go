package sys

import "sscmgroup.com/app/dto"

type MaterialForm struct {
	Id        int    `form:"id" json:"id"`
	Title     string `form:"title" json:"title"`
	Desc      string `form:"desc" json:"desc"`
	Thumbnail string `form:"thumbnail" json:"thumbnail"`
	Link      string `form:"link" json:"link"`
	Position  int    `form:"position" json:"position"`
	Status    int    `form:"status" json:"status"`
}

type MaterialSearchForm struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"`
	Position       int    `form:"position"`
	Status         int    `form:"status"`
	BeginTime      string `form:"beginTime"`
	EndTime        string `form:"endTime"`
}
