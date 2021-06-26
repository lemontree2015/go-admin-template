package sys

import (
	"errors"
	"github.com/sirupsen/logrus"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model/sys"
)

func (s *MgrService) AddSetting(form dto.SettingForm) (setting *sys.Setting, err error) {
	queryParams := make(map[string]interface{})
	queryParams["name = ?"] = form.Name
	setting, err = s.rep.GetSetting(queryParams)
	if setting != nil && setting.Id > 0 {
		logger.Logger.WithFields(logrus.Fields{
			"role_name": setting.Name,
		}).Info("该配制项已存在")
		return nil, errors.New("该配制项已存在")
	}

	setting, err = s.rep.CreateSetting(form)
	return
}

func (s *MgrService) GetSettingByPage(search *dto.SettingSearch) (count int, list []sys.Setting, err error) {
	count, list, err = s.rep.GetSettingByPage(search)
	return
}

func (s *MgrService) GetSettingById(Id int) (menu *sys.Setting, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	menu, err = s.rep.GetSetting(queryParams)
	if menu == nil || err != nil {
		err = errors.New("配制不存在")
		return
	}
	return
}

func (s *MgrService) UpdateSetting(row *sys.Setting, uf dto.SettingForm) (setting *sys.Setting, err error) {
	if uf.Name != "" && row.Name != uf.Name {
		row.Name = uf.Name
	}
	if uf.Title != "" && row.Title != uf.Title {
		row.Title = uf.Title
	}
	if uf.Content != "" && row.Content != uf.Content {
		row.Content = uf.Content
	}
	if uf.Status >= 0 && row.Status != uf.Status {
		row.Status = uf.Status
	}
	setting, err = s.rep.UpdateSetting(row)
	return
}
