package sys

import (
	"github.com/jinzhu/gorm"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/sys"
	"time"
)

func (r *Repository) GetSetting(qr map[string]interface{}) (row *sys.Setting, err error) {
	db := r.mgrDb
	row = &sys.Setting{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&row).Error
	if err == gorm.ErrRecordNotFound {
		row = nil
	}
	return
}

func (r *Repository) CreateSetting(f dto.SettingForm) (row *sys.Setting, err error) {
	now := time.Now()
	row = &sys.Setting{}
	row.Name = f.Name
	row.Title = f.Title
	row.Content = f.Content
	row.Status = 1
	row.Created = now
	row.Updated = now
	err = r.mgrDb.Create(&row).Error
	return
}

func (r *Repository) GetSettingByPage(rf *dto.SettingSearch) (count int, list []sys.Setting, err error) {
	db := r.mgrDb.Table("sys_settings")
	if rf.Name != "" {
		db = db.Where("name = ?", rf.Name)
	}
	if rf.Status >= 0 {
		db = db.Where("status = ?", rf.Status)
	}
	if rf.BeginTime != "" {
		db = db.Where("created >= ?", rf.BeginTime)
	}
	if rf.EndTime != "" {
		db = db.Where("created < ?", rf.EndTime)
	}

	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset((rf.GetPageIndex() - 1) * rf.GetPageSize()).Limit(rf.GetPageSize()).Find(&list).Error
	return
}

func (r *Repository) UpdateSetting(row *sys.Setting) (*sys.Setting, error) {
	err := r.mgrDb.Save(&row).Error
	return row, err
}
