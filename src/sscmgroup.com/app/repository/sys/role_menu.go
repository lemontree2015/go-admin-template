package sys

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/model/sys"
)

func (r *Repository) GetRoleMenu(roleId, menuId int) (rm *sys.RoleMenu, err error) {
	db := r.mgrDb.Table("sys_role_menu")
	rm = &sys.RoleMenu{}
	db = db.Where("role_id = ? and menu_id=?", roleId, menuId)
	err = db.First(&rm).Error
	if err == gorm.ErrRecordNotFound {
		rm = nil
	}
	return
}

func (r *Repository) SaveRoleMenu(roleId int, menuIds []int) (err error) {
	for _, menuID := range menuIds {
		roMenu, _ := r.GetRoleMenu(roleId, menuID)
		fmt.Println("SaveRoleMenu.........", roMenu, menuID)
		if roMenu != nil && roMenu.Id > 0 {
			roMenu.Status = 1
			err = r.mgrDb.Save(&roMenu).Error
		} else {
			roMenu = &sys.RoleMenu{}
			roMenu.MenuId = menuID
			roMenu.RoleId = roleId
			roMenu.Status = 1
			err = r.mgrDb.Create(&roMenu).Error
		}
	}
	return
}

func (r *Repository) RoleMenus(roleId int) (roleMenus []*sys.RoleMenu, err error) {
	db := r.mgrDb.Table("sys_role_menu").Where("role_id = ? and status=?", roleId, 1).Order("id ASC")
	err = db.Find(&roleMenus).Error
	return
}

func (r *Repository) DeleteRoleMenu(roleId int) (err error) {
	roleMenus, err := r.RoleMenus(roleId)
	if err == nil && len(roleMenus) > 0 {
		for _, item := range roleMenus {
			item.Status = 0
			err = r.mgrDb.Save(&item).Error
		}
	}
	return
}
