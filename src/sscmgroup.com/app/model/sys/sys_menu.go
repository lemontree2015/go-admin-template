package sys

const (
	TypeOfMenu   = 1
	TypeOfAction = 2
)

type Menu struct {
	Id         int    `gorm:"column:id;primary_key" json:"menuId"`
	ParentId   int    `gorm:"column:parent_id" json:"parentId"`
	Name       string `gorm:"column:name" json:"name"`
	Icon       string `gorm:"column:icon" json:"icon"`
	Path       string `gorm:"column:path" json:"path"`
	Action     string `gorm:"column:action" json:"action"`
	Component  string `gorm:"column:component" json:"component"`
	Sort       int    `gorm:"column:sort" json:"sort"`
	MenuType   int    `gorm:"column:menu_type" json:"menuType"`
	Method     string `gorm:"column:method" json:"method"`
	Permission string `gorm:"column:permission" json:"permission"`
	Status     int    `gorm:"column:status" json:"status"`
	Children   []Menu `json:"children" gorm:"-"`
}

func (m *Menu) TableName() string {
	return "sys_menu"
}

// ToRoleIDMap 转换为角色ID映射
func ToMenuIdMap(r []*Menu) map[int]*Menu {
	m := make(map[int]*Menu)
	for _, item := range r {
		m[item.Id] = item
	}
	return m
}
