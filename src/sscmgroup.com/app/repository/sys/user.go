package sys

import (
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/casbin"
	"sscmgroup.com/app/common"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/model/sys"
	"time"
)

func (r *Repository) GetUser(qr map[string]interface{}) (u *sys.User, err error) {
	db := r.mgrDb
	u = &sys.User{}
	for key, value := range qr {
		db = db.Where(key, value)
	}
	err = db.First(&u).Error
	if err == gorm.ErrRecordNotFound {
		u = nil
	}
	return
}

func (r *Repository) GetUserRoles(userId int) (roles []sys.Role, err error) {
	subQuery := r.mgrDb.Table("sys_user_role").
		Select("role_id").
		Where("user_id = ? and status = 1", userId).
		SubQuery()
	err = r.mgrDb.Table("sys_role").Where("id IN (?)", subQuery).Find(&roles).Error
	return
}

func (r *Repository) CreateUser(uf dto.UserForm) (u *sys.User, err error) {
	now := time.Now()
	u = &sys.User{}
	u.UserName = uf.UserName
	u.RealName = uf.RealName
	u.Password = common.SHA1String(uf.Password)
	u.Email = uf.Email
	u.Phone = uf.Phone
	u.Status = 1
	u.Created = now
	u.Updated = now

	err = r.mgrDb.Create(&u).Error
	if err != nil {
		return nil, err
	}
	err = r.SaveUserRole(u.Id, uf.RoleIds)
	if casbin.Enforcer != nil {
		casbin.LoadCasbinPolicy(casbin.Enforcer)
	}
	return
}

func (r *Repository) UpdateUser(user *sys.User, f dto.UserForm) (*sys.User, error) {
	err := r.mgrDb.Save(&user).Error
	if err != nil {
		return nil, err
	}
	if len(f.RoleIds) > 0 {
		//删除老菜单
		err = r.DeleteUserRole(user.Id)
		//增加新菜单
		err = r.SaveUserRole(user.Id, f.RoleIds)
	}

	if err != nil {
		return nil, err
	}
	if casbin.Enforcer != nil {
		casbin.LoadCasbinPolicy(casbin.Enforcer)
	}
	return user, err
}

func (r *Repository) UpdateUserInfo(user *sys.User) (err error) {
	err = r.mgrDb.Save(&user).Error
	return
}

func (r *Repository) GetSysUserPage(uf *dto.SysUserSearch) (count int, list []sys.User, err error) {
	db := r.mgrDb.Table("sys_user")
	if uf.UserId > 0 {
		db = db.Where("id = ?", uf.UserId)
	}
	if uf.Status >= 0 {
		db = db.Where("status = ?", uf.Status)
	}
	if uf.UserName != "" {
		db = db.Where("user_name = ?", uf.UserName)
	}
	if uf.Email != "" {
		db = db.Where("email = ?", uf.Email)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset((uf.GetPageIndex() - 1) * uf.GetPageSize()).Limit(uf.GetPageSize()).Find(&list).Error
	return
}
