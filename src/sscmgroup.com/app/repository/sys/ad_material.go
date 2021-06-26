package sys

import (
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/cms"
	"time"
)

func (r *Repository) GetMaterial(qr map[string]interface{}) (row *cms.AdMaterials, err error) {
	db := r.cmsDb
	row = &cms.AdMaterials{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&row).Error
	return
}

func (r *Repository) CreateMaterial(form dto.MaterialForm) (row *cms.AdMaterials, err error) {
	now := time.Now()
	row = &cms.AdMaterials{}
	row.Title = form.Title
	row.Desc = form.Desc
	row.Thumbnail = form.Thumbnail
	row.Link = form.Link
	row.Position = form.Position
	row.Status = 1
	row.Created = now
	row.Updated = now
	err = r.cmsDb.Create(&row).Error
	return
}

func (r *Repository) GetMaterialByPage(form *dto.MaterialSearchForm) (count int, list []cms.AdMaterials, err error) {
	db := r.cmsDb.Table("ad_materials")
	if form.Title != "" {
		db = db.Where("title = ?", form.Title)
	}
	if form.Position >= 0 {
		db = db.Where("position = ?", form.Position)
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

func (r *Repository) UpdateMaterial(row *cms.AdMaterials) (*cms.AdMaterials, error) {
	err := r.cmsDb.Save(&row).Error
	return row, err
}
