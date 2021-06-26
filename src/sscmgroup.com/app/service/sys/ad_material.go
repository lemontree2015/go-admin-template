package sys

import (
	"errors"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/cms"
	"time"
)

func (s *MgrService) GetMaterialByPage(rf *dto.MaterialSearchForm) (count int, list []cms.AdMaterials, err error) {
	count, list, err = s.rep.GetMaterialByPage(rf)
	return
}

func (s *MgrService) GetMaterialById(Id int) (row *cms.AdMaterials, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	row, err = s.rep.GetMaterial(queryParams)
	return
}

func (s *MgrService) AddMaterial(form dto.MaterialForm) (row *cms.AdMaterials, err error) {
	queryParams := make(map[string]interface{})
	queryParams["title = ?"] = form.Title
	row, err = s.rep.GetMaterial(queryParams)
	if row != nil && row.Id > 0 {
		return nil, errors.New("该新闻已存在")
	}
	row, err = s.rep.CreateMaterial(form)
	return
}

func (s *MgrService) UpdateMaterial(m *cms.AdMaterials, f dto.MaterialForm) (row *cms.AdMaterials, err error) {
	now := time.Now()
	m.Title = f.Title
	m.Desc = f.Desc
	m.Thumbnail = f.Thumbnail
	m.Position = f.Position
	m.Link = f.Link
	m.Status = f.Status
	m.Updated = now
	row, err = s.rep.UpdateMaterial(m)
	return
}
