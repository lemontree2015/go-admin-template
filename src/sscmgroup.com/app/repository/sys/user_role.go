package sys

import (
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/model/sys"
)

func (r *Repository) GetOneUserRole(userId, roleId int) (ur *sys.UserRole, err error) {
	db := r.mgrDb
	ur = &sys.UserRole{}
	db = db.Where("user_id = ? and role_id = ?", userId, roleId)
	err = db.First(&ur).Error
	if err == gorm.ErrRecordNotFound {
		ur = nil
	}
	return
}

func (r *Repository) UserRoles(userId int) (urs []*sys.UserRole, err error) {
	db := r.mgrDb.Table("sys_user_role").Where("user_id = ? and status=?", userId, 1).Order("id ASC")
	err = db.Find(&urs).Error
	return
}

func (r *Repository) SaveUserRole(userId int, roleIds []int) (err error) {
	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			ur, _ := r.GetOneUserRole(userId, roleId)
			if ur != nil && ur.Id > 0 {
				ur.Status = 1
				err = r.mgrDb.Save(ur).Error
			} else {
				ur = &sys.UserRole{}
				ur.RoleId = roleId
				ur.UserId = userId
				ur.Status = 1
				err = r.mgrDb.Create(ur).Error
			}
		}
	}
	return
}

func (r *Repository) DeleteUserRole(userId int) (err error) {
	urs, err := r.UserRoles(userId)
	if err == nil && len(urs) > 0 {
		for _, item := range urs {
			item.Status = 0
			err = r.mgrDb.Save(&item).Error
		}
	}
	return
}
