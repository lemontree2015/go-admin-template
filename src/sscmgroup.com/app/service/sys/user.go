package sys

import (
	"errors"
	"fmt"
	"sscmgroup.com/app/common"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/sys"
	"time"
)

func (s *MgrService) GetUserById(Id int) (user *sys.User, err error) {
	queryParams := make(map[string]interface{})
	queryParams["id = ?"] = Id
	user, err = s.rep.GetUser(queryParams)
	if user == nil || err != nil {
		err = errors.New("用户不存在")
		return
	}
	rs, _ := s.rep.GetUserRoles(user.Id)
	var roleIds []int
	for _, r := range rs {
		roleIds = append(roleIds, r.Id)
	}
	permissions, _ := s.rep.GetRolesPermissions(roleIds)
	user.Permissions = permissions
	user.Roles = rs
	return
}

func (s *MgrService) Login(lf dto.LoginForm, loginIp string) (user *sys.User, err error) {
	queryParams := make(map[string]interface{})
	queryParams["user_name = ?"] = lf.UserName
	user, err = s.rep.GetUser(queryParams)
	if user == nil || err != nil {
		err = errors.New("用户不存在")
		return
	}
	if user.Password != common.SHA1String(lf.Password) {
		err = errors.New("密码错误")
		return
	}

	user.LastLoginIp = loginIp
	user.LastLoginTime = time.Now()
	s.rep.UpdateUserInfo(user)

	roles, _ := s.rep.GetUserRoles(user.Id)
	fmt.Println("service roles....", roles)
	if len(roles) > 0 {
		user.Roles = roles
	}
	return
}

func (s *MgrService) AddUser(uf dto.UserForm) (user *sys.User, err error) {
	queryParams := make(map[string]interface{})
	queryParams["user_name = ?"] = uf.UserName
	user, err = s.rep.GetUser(queryParams)
	if user != nil && user.Id > 0 {
		err = errors.New("该用户名已存在")
		return
	}

	user, err = s.rep.CreateUser(uf)
	return
}

func (s *MgrService) UpdateUser(u *sys.User, uf dto.UserForm) (user *sys.User, err error) {
	now := time.Now()
	if uf.UserName != "" && u.UserName != uf.UserName {
		u.UserName = uf.UserName
	}
	if uf.RealName != "" && u.RealName != uf.RealName {
		u.RealName = uf.RealName
	}
	if uf.Email != "" && u.Email != uf.Email {
		u.Email = uf.Email
	}
	if uf.Phone != "" && u.Phone != uf.Phone {
		u.Phone = uf.Phone
	}
	if uf.Status >= 0 && u.Status != uf.Status {
		u.Status = uf.Status
	}
	u.Updated = now
	user, err = s.rep.UpdateUser(u, uf)
	return
}

func (s *MgrService) UpdateUserInfo(u *sys.User) (err error) {
	err = s.rep.UpdateUserInfo(u)
	return
}

func (s *MgrService) ChangeUserPwd(u *sys.User, oldPwd, newPwd string) (err error) {
	oldHash := common.SHA1String(oldPwd)
	if u.Password != oldHash {
		err = errors.New("旧密码错误")
		return
	}
	newHash := common.SHA1String(newPwd)
	u.Password = newHash
	err = s.rep.UpdateUserInfo(u)
	return
}

func (s *MgrService) GetSysUserPage(uf *dto.SysUserSearch) (count int, list []sys.User, err error) {
	count, list, err = s.rep.GetSysUserPage(uf)
	return
}
