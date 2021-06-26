package sys

type RoleMenu struct {
	Id     int `gorm:"column:id;primary_key" json:"id"`
	RoleId int `gorm:"column:role_id" json:"roleId"`
	MenuId int `gorm:"column:menu_id" json:"menuId"`
	Status int `gorm:"column:status" json:"status"`
}

func (rm *RoleMenu) TableName() string {
	return "sys_role_menu"
}

// RoleMenus 角色菜单列表
type RoleMenus []*RoleMenu

// ToRoleIDMap 转换为角色ID 资源映射
func ToRoleIDResourceMap(rr []*RoleMenu) map[int]RoleMenus {
	m := make(map[int]RoleMenus)
	for _, item := range rr {
		m[item.RoleId] = append(m[item.RoleId], item)
	}
	return m
}
