package sys

import (
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/casbin"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/sys"
	"time"
)

func (r *Repository) GetRole(qr map[string]interface{}) (role *sys.Role, err error) {
	db := r.mgrDb
	role = &sys.Role{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&role).Error
	if err == gorm.ErrRecordNotFound {
		role = nil
	}
	return
}

func (r *Repository) CreateRole(f dto.RoleForm) (role *sys.Role, err error) {
	now := time.Now()
	role = &sys.Role{}
	role.Title = f.Title
	role.Name = f.Name
	role.Remark = f.Remark
	role.Status = 1
	role.Created = now
	role.Updated = now

	err = r.mgrDb.Create(&role).Error
	if err != nil {
		return nil, err
	}
	err = r.SaveRoleMenu(role.Id, f.MenuIds)
	if err != nil {
		return nil, err
	}
	if casbin.Enforcer != nil {
		casbin.LoadCasbinPolicy(casbin.Enforcer)
	}
	return role, err
}

func (r *Repository) UpdateRole(role *sys.Role, f dto.RoleForm) (*sys.Role, error) {
	err := r.mgrDb.Save(&role).Error
	if err != nil {
		return nil, err
	}
	err = r.DeleteRoleMenu(role.Id)
	if len(f.MenuIds) > 0 {
		//增加新菜单
		err = r.SaveRoleMenu(role.Id, f.MenuIds)
	}

	if err != nil {
		return nil, err
	}
	if casbin.Enforcer != nil {
		casbin.LoadCasbinPolicy(casbin.Enforcer)
	}
	return role, err
}

func (r *Repository) GetSysRolePage(rf *dto.SysRoleSearch) (count int, list []sys.Role, err error) {
	db := r.mgrDb.Table("sys_role")
	if rf.RoleId > 0 {
		db = db.Where("id = ?", rf.RoleId)
	}
	if rf.Status != "-1" {
		db = db.Where("status = ?", rf.Status)
	}
	if rf.BeginTime != "" {
		db = db.Where("created >= ?", rf.BeginTime)
	}
	if rf.EndTime != "" {
		db = db.Where("created < ?", rf.EndTime)
	}
	if rf.RoleName != "" {
		db = db.Where("name = ?", rf.RoleName)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset((rf.GetPageIndex() - 1) * rf.GetPageSize()).Limit(rf.GetPageSize()).Find(&list).Error
	return
}

func (r *Repository) GetAllSysRoles() (list []sys.Role, err error) {
	db := r.mgrDb.Table("sys_role").Where("status=?", 1).Order("id ASC")
	err = db.Find(&list).Error
	return
}
