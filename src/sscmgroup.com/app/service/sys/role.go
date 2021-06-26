package sys

import (
	"errors"
	"github.com/sirupsen/logrus"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model/sys"
	"time"
)

func (s *MgrService) AddRole(f dto.RoleForm) (role *sys.Role, err error) {
	queryParams := make(map[string]interface{})
	queryParams["name = ?"] = f.Name
	role, err = s.rep.GetRole(queryParams)
	if role != nil && role.Id > 0 {
		logger.Logger.WithFields(logrus.Fields{
			"role_name":  role.Name,
			"role_title": role.Title,
		}).Info("该角色已存在")
		return nil, errors.New("该角色已存在")
	}

	role, err = s.rep.CreateRole(f)
	return
}

func (s *MgrService) UpdateRole(r *sys.Role, f dto.RoleForm) (role *sys.Role, err error) {
	now := time.Now()
	r.Name = f.Name
	r.Title = f.Title
	r.Remark = f.Remark
	r.Status = f.Status
	r.Updated = now
	role, err = s.rep.UpdateRole(r, f)
	return
}

func (s *MgrService) GetRoleById(Id int) (role *sys.Role, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	role, err = s.rep.GetRole(queryParams)
	if role == nil || err != nil {
		err = errors.New("角色不存在")
		return
	}

	//过滤掉顶级菜单，否则vue菜单会全选中
	menuIds, err := s.GetRoleSubMenuIds(Id)
	role.MenuIds = menuIds
	return
}

func (s *MgrService) GetRoleSubMenuIds(roleId int) (m []int, err error) {
	m, err = s.rep.GetRoleSubMenuIds(roleId)
	return
}

func (s *MgrService) GetSysRolePage(rf *dto.SysRoleSearch) (count int, list []sys.Role, err error) {
	count, list, err = s.rep.GetSysRolePage(rf)
	return
}

func (s *MgrService) GetAllSysRoles() (list []sys.Role, err error) {
	list, err = s.rep.GetAllSysRoles()
	return
}
