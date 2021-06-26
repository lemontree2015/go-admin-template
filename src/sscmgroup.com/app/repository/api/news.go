package api

import (
	"sscmgroup.com/app/model/cms"
)

const (
	pageSize = 10
)

func (r *Repository) GetSubCateByPid(lang string, pId int) (list []*cms.NewsCategory, err error) {
	db := r.mgrDb.Table("news_category")
	db = db.Where("parent_id = ? and lang = ? and status = 1", pId, lang)
	err = db.Find(&list).Error
	return
}

func (r *Repository) GetCateNewsByPid(cateId int) (list []cms.News, err error) {
	db := r.mgrDb.Table("news_titles")
	db = db.Where("cate_id = ?  and status = 1", cateId)
	err = db.Find(&list).Error
	return
}

func (r *Repository) GetNewsDetail(id int) (row *cms.News, err error) {
	row = &cms.News{}
	db := r.mgrDb.Table("news_titles")
	db = db.Where("id = ?", id)
	err = db.First(&row).Error
	if err != nil {
		return
	}
	content := &cms.NewsContent{}
	db = r.mgrDb.Table("news_contents")
	db = db.Where("news_id = ?", id)
	err = db.First(&content).Error
	if err == nil && content != nil {
		row.Content = content.Content
	}
	return
}

func (r *Repository) GetCateNews(cateId, page int) (rows []cms.News, count int, err error) {
	db := r.cmsDb

	db = db.Select("a.*,b.content").Table("news_titles a").Joins("left join news_contents b on a.id = b.news_id").Where("cate_id = ? and a.status = 1", cateId)
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&rows).Error
	err = db.Find(&rows).Error
	return
}

func (r *Repository) GetAboutUsNews(lang string, cId int) (rows []cms.News, err error) {
	db := r.mgrDb.Table("news_titles")
	db = db.Where("cate_id = ? and lang = ? and status = 1", cId, lang)
	err = db.Find(&rows).Error
	return
}

func (r *Repository) GetAdListByPosition(pId int) (list []cms.AdMaterials, err error) {
	db := r.cmsDb.Table("ad_materials")
	db = db.Where("position = ? and status = 1", pId)
	err = db.Find(&list).Error
	return
}
