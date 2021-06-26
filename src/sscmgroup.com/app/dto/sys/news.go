package sys

import "sscmgroup.com/app/dto"

type NewsCategoryForm struct {
	CateId   int    `form:"cateId" json:"cateId"`
	ParentId int    `form:"parentId" json:"parentId"`
	Name     string `form:"name" json:"name"`
	Status   int    `form:"status" json:"status"`
}

type CategorySearchForm struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"`
	Status         int    `form:"status"`
	BeginTime      string `form:"beginTime"`
	EndTime        string `form:"endTime"`
}

type NewsSearchForm struct {
	dto.Pagination `search:"-"`
	CateId         int    `form:"cateId"`
	Title          string `form:"title"`
	Status         int    `form:"status"`
	BeginTime      string `form:"beginTime"`
	EndTime        string `form:"endTime"`
}

type NewsForm struct {
	Id  int    `form:"id" json:"id"`
	CateId  int    `form:"cateId" json:"cateId"`
	Title   string `form:"title" json:"title"`
	Desc    string `form:"desc" json:"desc"`
	Thumbnail    string `form:"thumbnail" json:"thumbnail"`
	NewsType  int    `form:"news_type" json:"newsType"`
	Content string `form:"content" json:"content"`
	Status  int    `form:"status" json:"status"`
}
