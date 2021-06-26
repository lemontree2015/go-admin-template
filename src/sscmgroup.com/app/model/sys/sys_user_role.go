package sys

type UserRole struct {
	Id     int `gorm:"column:id;primary_key" json:"id"`
	UserId int `gorm:"column:user_id" json:"userId"`
	RoleId int `gorm:"column:role_id" json:"roleId"`
	Status int `gorm:"column:status" json:"status"`
}

func (r *UserRole) TableName() string {
	return "sys_user_role"
}

// RoleMenus 角色菜单列表
type UserRoleIds []int

// ToRoleIDMap 转换为角色ID映射
func ToUserRoleIDMap(ur []*UserRole) map[int]UserRoleIds {
	m := make(map[int]UserRoleIds)
	for _, item := range ur {
		m[item.UserId] = append(m[item.UserId], item.RoleId)
	}
	return m
}
