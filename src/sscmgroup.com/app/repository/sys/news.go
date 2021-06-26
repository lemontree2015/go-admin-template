package sys

import (
	"github.com/jinzhu/gorm"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/cms"
	"time"
)

func (r *Repository) GetNewsCategory(qr map[string]interface{}) (row *cms.NewsCategory, err error) {
	db := r.cmsDb
	row = &cms.NewsCategory{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&row).Error
	if err == gorm.ErrRecordNotFound {
		row = nil
	}
	return
}

func (r *Repository) CreateNewsCategory(form dto.NewsCategoryForm) (row *cms.NewsCategory, err error) {
	now := time.Now()
	row = &cms.NewsCategory{}
	row.ParentId = form.ParentId
	row.Name = form.Name
	row.Status = 1
	row.Created = now
	row.Updated = now
	err = r.cmsDb.Create(&row).Error
	return
}

func (r *Repository) UpdateCategory(cate *cms.NewsCategory) (*cms.NewsCategory, error) {
	err := r.cmsDb.Save(&cate).Error
	return cate, err
}

func (r *Repository) GetCategoryByPage(form *dto.CategorySearchForm) (count int, list []cms.NewsCategory, err error) {
	db := r.cmsDb.Table("news_category")
	if form.Name != "" {
		db = db.Where("name = ?", form.Name)
	}
	if form.Status >= 0 {
		db = db.Where("status = ?", form.Status)
	}
	if form.BeginTime != "" {
		db = db.Where("created >= ?", form.BeginTime)
	}
	if form.EndTime != "" {
		db = db.Where("created < ?", form.EndTime)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset((form.GetPageIndex() - 1) * form.GetPageSize()).Limit(form.GetPageSize()).Find(&list).Error
	return
}

func (r *Repository) GetAllCategories() (list []cms.NewsCategory, err error) {
	db := r.cmsDb.Table("news_category")
	db = db.Where("status = ?", 1)
	err = db.Find(&list).Error
	return
}

func (r *Repository) GetNewsByPage(form *dto.NewsSearchForm) (count int, list []cms.News, err error) {
	db := r.cmsDb.Table("news_titles")
	if form.CateId > 0 {
		db = db.Where("cate_id = ?", form.CateId)
	}
	if form.Title != "" {
		db = db.Where("title = ?", form.Title)
	}
	if form.Status >= 0 {
		db = db.Where("status = ?", form.Status)
	}
	if form.BeginTime != "" {
		db = db.Where("created >= ?", form.BeginTime)
	}
	if form.EndTime != "" {
		db = db.Where("created < ?", form.EndTime)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset((form.GetPageIndex() - 1) * form.GetPageSize()).Limit(form.GetPageSize()).Find(&list).Error
	return
}

func (r *Repository) GetNews(qr map[string]interface{}) (row *cms.News, err error) {
	db := r.cmsDb
	row = &cms.News{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&row).Error
	if err == gorm.ErrRecordNotFound {
		row = nil
		return
	}
	content := &cms.NewsContent{}
	db = db.Where("news_id = ?", row.Id)
	err = db.First(&content).Error
	if err == nil {
		row.Content = content.Content
	}
	return
}

func (r *Repository) CreateNews(form dto.NewsForm) (row *cms.News, err error) {
	now := time.Now()
	row = &cms.News{}
	row.CateId = form.CateId
	row.Title = form.Title
	row.Desc = form.Desc
	row.Thumbnail = form.Thumbnail
	row.NewsType = form.NewsType
	row.Status = 1
	row.Created = now
	row.Updated = now
	err = r.cmsDb.Create(&row).Error
	if err != nil {
		return
	}

	content := &cms.NewsContent{}
	content.NewsId = row.Id
	content.Content = form.Content
	content.Created = now
	content.Updated = now
	err = r.cmsDb.Create(&content).Error
	if err != nil {
		row.Content = form.Content
	}
	return
}

func (r *Repository) UpdateNews(news *cms.News) (*cms.News, error) {
	err := r.cmsDb.Save(&news).Error
	if err != nil {
		return nil, err
	}
	content := &cms.NewsContent{}
	err = r.cmsDb.Where("news_id = ?", news.Id).First(&content).Error
	if err == nil {
		content.Content = news.Content
		err = r.cmsDb.Save(&content).Error
	}
	return news, err
}
