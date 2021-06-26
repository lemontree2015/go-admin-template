package sys

import (
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/common"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/sys"
)

func (r *Repository) GetMenu(qr map[string]interface{}) (menu *sys.Menu, err error) {
	db := r.mgrDb
	menu = &sys.Menu{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&menu).Error
	if err == gorm.ErrRecordNotFound {
		menu = nil
	}
	return
}

func (r *Repository) CreateMenu(f dto.MenuForm) (m *sys.Menu, err error) {
	m = &sys.Menu{}
	m.ParentId = f.ParentId
	m.Name = f.Name
	m.Icon = f.Icon
	m.Path = f.Path
	m.Action = f.Action
	m.Component = f.Component
	m.Sort = f.Sort
	m.MenuType = f.MenuType
	m.Method = f.Method
	m.Permission = f.Permission
	m.Status = 1
	err = r.mgrDb.Create(&m).Error
	return
}

func (r *Repository) UpdateMenu(menu *sys.Menu) (*sys.Menu, error) {
	err := r.mgrDb.Save(&menu).Error
	return menu, err
}

func (r *Repository) GetAllSysMenus() (menus []sys.Menu, err error) {
	db := r.mgrDb.Table("sys_menu").Where("status=?", 1).Order("sort ASC")
	err = db.Find(&menus).Error
	return
}

func (r *Repository) GetSuperRoleMenus() (menus []sys.Menu, err error) {
	db := r.mgrDb.Table("sys_menu").Where("status=? and menu_type = ?", 1, sys.TypeOfMenu).Order("sort ASC")
	err = db.Find(&menus).Error
	return
}

func (r *Repository) GetRoleAccessSideBar(roleIds []int) (menus []sys.Menu, err error) {
	subQuery := r.mgrDb.Table("sys_role_menu").
		Select("distinct sys_role_menu.menu_id").
		Where("role_id in(?) and sys_role_menu.status = ? and menu_type = ?", roleIds, 1, sys.TypeOfMenu).
		SubQuery()
	err = r.mgrDb.Table("sys_menu").Where("id IN (?) and sys_menu.status = 1", subQuery).Find(&menus).Error
	return
}

func (r *Repository) GetRoleSubMenuIds(roleId int) (menuIds []int, err error) {
	var menus []sys.Menu
	subQuery := r.mgrDb.Table("sys_role_menu").
		Select("distinct sys_role_menu.menu_id").
		Where("role_id = ? and sys_role_menu.status = ?", roleId, 1).
		Where(" sys_role_menu.menu_id not in(select sys_menu.parent_id from sys_role_menu LEFT JOIN sys_menu on sys_menu.id=sys_role_menu.menu_id where role_id = ? and sys_role_menu.status = ? and parent_id is not null)", roleId, 1).
		SubQuery()
	err = r.mgrDb.Table("sys_menu").Where("id IN (?)", subQuery).Find(&menus).Error
	for _, m := range menus {
		menuIds = append(menuIds, m.Id)
	}
	return
}

func (r *Repository) GetRolesPermissions(roleIds []int) (permissions []string, err error) {
	if len(roleIds) > 0 {
		if common.InArray(common.AdminRoleId, roleIds) {
			permissions = append(permissions, "*:*:*")
		} else {
			var menus []sys.Menu
			subQuery := r.mgrDb.Table("sys_role_menu").
				Select("distinct sys_role_menu.menu_id").
				Where("role_id in(?) and sys_role_menu.status = ?", roleIds, 1).
				Where("sys_role_menu.menu_id in(select sys_menu.id from sys_menu where sys_menu.status = ? and permission is not null and menu_type = ?)", 1, 2).
				SubQuery()
			err = r.mgrDb.Table("sys_menu").Where("id IN (?)", subQuery).Find(&menus).Error
			if err == nil {
				for _, menu := range menus {
					permissions = append(permissions, menu.Permission)
				}
			}
		}
	}

	return
}
