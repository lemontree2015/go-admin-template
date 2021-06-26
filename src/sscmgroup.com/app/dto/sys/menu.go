package sys

type MenuForm struct {
	MenuId     int    `form:"menuId" json:"menuId"`
	ParentId   int    `form:"parentId" json:"parentId"`
	Name       string `form:"name" json:"name" binding:"required"`
	Icon       string `form:"icon" json:"icon"`
	Path       string `form:"path" json:"path"`
	Action     string `form:"action" json:"action"`
	Component  string `form:"component" json:"component"`
	Sort       int    `form:"sort" json:"sort"`
	MenuType   int    `form:"menuType" json:"menuType"`
	Method     string `form:"method" json:"method"`
	Permission string `form:"permission" json:"permission"`
	Status     int    `form:"status" json:"status"`
}

type MenuLabel struct {
	Id       int         `json:"id" gorm:"-"`
	Label    string      `json:"label" gorm:"-"`
	Children []MenuLabel `json:"children" gorm:"-"`
}
