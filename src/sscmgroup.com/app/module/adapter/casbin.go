package adapter

import (
	"context"
	"fmt"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model/sys"
)

// CasbinAdapter casbin适配器
type CasbinAdapter struct {
	MgrDb *gorm.DB
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(model casbinModel.Model) error {
	logger.Logger.Debug("load Casbin Policy")
	ctx := context.Background()
	err := a.loadRolePolicy(ctx, model)
	if err != nil {
		logger.Logger.Error("Load casbin role policy error: %s", err.Error())
		return err
	}

	err = a.loadUserPolicy(ctx, model)
	if err != nil {
		logger.Logger.Error("Load casbin user policy error: %s", err.Error())
		return err
	}
	return nil
}

// 加载角色策略(p,role_id,path,method)
func (a *CasbinAdapter) loadRolePolicy(ctx context.Context, m casbinModel.Model) error {
	var roles []*sys.Role
	err := a.MgrDb.Model(&sys.Role{}).Where("status = 1").Scan(&roles).Error
	if err != nil {
		return err
	} else if len(roles) == 0 {
		return nil
	}
	var roleResource []*sys.RoleMenu
	err = a.MgrDb.Model(&sys.RoleMenu{}).Where("status = 1").Scan(&roleResource).Error
	if err != nil {
		return err
	}
	mRoleResource := sys.ToRoleIDResourceMap(roleResource)

	var resources []*sys.Menu
	err = a.MgrDb.Model(&sys.Menu{}).Where("status = 1").Scan(&resources).Error
	if err != nil {
		return err
	}
	mResource := sys.ToMenuIdMap(resources)

	for _, role := range roles {
		mCache := make(map[string]struct{})
		if rrs, ok := mRoleResource[role.Id]; ok {
			for _, rr := range rrs {
				if r, ok := mResource[rr.MenuId]; ok {
					if r.Action == "" || r.Method == "" {
						continue
					} else if _, ok := mCache[r.Action+r.Method]; ok {
						continue
					}
					mCache[r.Action+r.Method] = struct{}{}
					line := fmt.Sprintf("p,%s,%s,%s", role.Name, r.Action, r.Method)
					persist.LoadPolicyLine(line, m)
				}
			}
		}
	}
	return nil
}

// 加载用户策略(g,user_id,role_id)
func (a *CasbinAdapter) loadUserPolicy(ctx context.Context, m casbinModel.Model) error {
	var roles []*sys.Role
	err := a.MgrDb.Model(&sys.Role{}).Where("status = 1").Scan(&roles).Error
	if err != nil {
		return err
	} else if len(roles) == 0 {
		return nil
	}
	mRoles := sys.ToRoleIDMap(roles)

	var users []*sys.User
	err = a.MgrDb.Model(&sys.User{}).Where("status = 1").Scan(&users).Error
	if err != nil {
		return err
	} else if len(users) > 0 {
		var userRoles []*sys.UserRole
		err = a.MgrDb.Model(&sys.UserRole{}).Where("status = 1").Scan(&userRoles).Error
		if err != nil {
			return err
		}
		mUserRoles := sys.ToUserRoleIDMap(userRoles)
		for _, u := range users {
			if ur, ok := mUserRoles[u.Id]; ok {
				for _, r := range ur {
					if ri, ok := mRoles[r]; ok {
						line := fmt.Sprintf("g,%s,%s", u.UserName, ri.Name)
						persist.LoadPolicyLine(line, m)
					}
				}
			}
		}
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
