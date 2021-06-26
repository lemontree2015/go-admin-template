package sys

import (
	"errors"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/cms"
	"time"
)

func (s *MgrService) GetCategoryByPage(rf *dto.CategorySearchForm) (count int, list []cms.NewsCategory, err error) {
	count, list, err = s.rep.GetCategoryByPage(rf)
	return
}

func (s *MgrService) GetAllCategories() (rows []cms.NewsCategory, err error) {
	list, err := s.rep.GetAllCategories()
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		menusInfo := makeCategoryTree(&list, list[i])
		rows = append(rows, menusInfo)
	}
	return
}

func makeCategoryTree(rows *[]cms.NewsCategory, row cms.NewsCategory) cms.NewsCategory {
	list := *rows

	min := make([]cms.NewsCategory, 0)
	for j := 0; j < len(list); j++ {

		if row.Id != list[j].ParentId {
			continue
		}
		mi := cms.NewsCategory{}
		mi.Id = list[j].Id
		mi.Name = list[j].Name
		mi.ParentId = list[j].ParentId
		mi.Children = []cms.NewsCategory{}
		min = append(min, mi)
	}
	row.Children = min
	return row
}

func (s *MgrService) GetCategoryById(Id int) (row *cms.NewsCategory, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	row, err = s.rep.GetNewsCategory(queryParams)
	return
}

func (s *MgrService) AddCategory(form dto.NewsCategoryForm) (row *cms.NewsCategory, err error) {
	queryParams := make(map[string]interface{})
	queryParams["name = ?"] = form.Name
	row, err = s.rep.GetNewsCategory(queryParams)
	if row != nil && row.Id > 0 {
		return nil, errors.New("该分类已存在")
	}
	row, err = s.rep.CreateNewsCategory(form)
	return
}

func (s *MgrService) UpdateCategory(cate *cms.NewsCategory, f dto.NewsCategoryForm) (row *cms.NewsCategory, err error) {
	now := time.Now()
	cate.ParentId = f.ParentId
	cate.Name = f.Name
	cate.Status = f.Status
	cate.Updated = now
	row, err = s.rep.UpdateCategory(cate)
	return
}

func (s *MgrService) GetNewsById(Id int) (row *cms.News, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	row, err = s.rep.GetNews(queryParams)
	return
}

func (s *MgrService) GetNewsByPage(rf *dto.NewsSearchForm) (count int, list []cms.News, err error) {
	count, list, err = s.rep.GetNewsByPage(rf)
	return
}

func (s *MgrService) AddNews(form dto.NewsForm) (row *cms.News, err error) {
	queryParams := make(map[string]interface{})
	queryParams["title = ?"] = form.Title
	queryParams["cate_id = ?"] = form.CateId
	row, err = s.rep.GetNews(queryParams)
	if row != nil && row.Id > 0 {
		return nil, errors.New("该新闻已存在")
	}
	row, err = s.rep.CreateNews(form)
	return
}

func (s *MgrService) UpdateNews(news *cms.News, f dto.NewsForm) (row *cms.News, err error) {
	now := time.Now()
	news.CateId = f.CateId
	news.Title = f.Title
	news.Desc = f.Desc
	news.Thumbnail = f.Thumbnail
	news.NewsType = f.NewsType
	news.Content = f.Content
	news.Status = f.Status
	news.Updated = now
	row, err = s.rep.UpdateNews(news)
	return
}
