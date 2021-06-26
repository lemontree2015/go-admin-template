package sys

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sscmgroup.com/app/common"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model/sys"
)

func (s *MgrService) GetMenuById(Id int) (menu *sys.Menu, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	menu, err = s.rep.GetMenu(queryParams)
	if menu == nil || err != nil {
		err = errors.New("菜单不存在")
		return
	}
	return
}

func (s *MgrService) AddMenu(f dto.MenuForm) (menu *sys.Menu, err error) {
	queryParams := make(map[string]interface{})
	queryParams["name = ?"] = f.Name
	menu, err = s.rep.GetMenu(queryParams)
	if menu != nil && menu.Id > 0 {
		logger.Logger.WithFields(logrus.Fields{
			"menu_name": menu.Name,
		}).Info("该菜单已存在")
		return nil, errors.New("该菜单已存在")
	}
	menu, err = s.rep.CreateMenu(f)
	return
}

func (s *MgrService) UpdateMenu(m *sys.Menu, uf dto.MenuForm) (menu *sys.Menu, err error) {
	if uf.ParentId >= 0 && m.ParentId != uf.ParentId {
		m.ParentId = uf.ParentId
	}
	if uf.Name != "" && m.Name != uf.Name {
		m.Name = uf.Name
	}
	if uf.Icon != "" && m.Icon != uf.Icon {
		m.Icon = uf.Icon
	}
	if uf.Path != "" && m.Path != uf.Path {
		m.Path = uf.Path
	}
	if uf.Action != "" && m.Action != uf.Action {
		m.Action = uf.Action
	}
	if uf.Component != "" && m.Component != uf.Component {
		m.Component = uf.Component
	}
	if uf.Sort >= 0 && m.Sort != uf.Sort {
		m.Sort = uf.Sort
	}
	if uf.MenuType > 0 && m.MenuType != uf.MenuType {
		m.MenuType = uf.MenuType
	}
	if uf.Method != "" && m.Method != uf.Method {
		m.Method = uf.Method
	}
	if uf.Permission != "" && m.Permission != uf.Permission {
		m.Permission = uf.Permission
	}
	if uf.Status >= 0 && m.Status != uf.Status {
		m.Status = uf.Status
	}
	menu, err = s.rep.UpdateMenu(m)
	return
}

func (s *MgrService) GetUserAccessSideBar(userId int) (m []sys.Menu, err error) {
	rs, err := s.rep.GetUserRoles(userId)
	roleIds := sys.ToRoleIDs(rs)
	if len(roleIds) > 0 {
		fmt.Println("roleIds...", roleIds)
		menus := make([]sys.Menu, 0)
		if common.InArray(common.AdminRoleId, roleIds) {
			menus, err = s.rep.GetSuperRoleMenus()
		} else {
			menus, err = s.rep.GetRoleAccessSideBar(roleIds)
		}
		for i := 0; i < len(menus); i++ {
			if menus[i].ParentId != 0 {
				continue
			}
			menusInfo := makeMenuTree(&menus, menus[i])
			m = append(m, menusInfo)
		}
	}
	return
}

func (s *MgrService) GetAllSysMenuLabels() (m []dto.MenuLabel, err error) {
	list, err := s.rep.GetAllSysMenus()
	m = make([]dto.MenuLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = list[i].Id
		e.Label = list[i].Name
		node := makeMenuLabelTree(&list, e)
		m = append(m, node)
	}
	return
}

func (s *MgrService) GetAllSysMenus() (m []sys.Menu, err error) {
	menus, err := s.rep.GetAllSysMenus()
	//m = make([]sys.Menu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := makeMenuTree(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return
}

// menuLabelCall 递归构造组织数据
func makeMenuLabelTree(menus *[]sys.Menu, dept dto.MenuLabel) dto.MenuLabel {
	list := *menus

	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].Id
		mi.Label = list[j].Name
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType == 1 {
			//菜单
			ms := makeMenuLabelTree(menus, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

func makeMenuTree(menus *[]sys.Menu, menu sys.Menu) sys.Menu {
	list := *menus

	min := make([]sys.Menu, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].ParentId {
			continue
		}
		mi := sys.Menu{}
		mi.Id = list[j].Id
		mi.Name = list[j].Name
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.Action = list[j].Action
		mi.MenuType = list[j].MenuType
		mi.Method = list[j].Method
		mi.ParentId = list[j].ParentId
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Permission = list[j].Permission
		mi.Children = []sys.Menu{}

		if mi.MenuType == 1 {
			ms := makeMenuTree(menus, mi)
			min = append(min, ms)

		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}
