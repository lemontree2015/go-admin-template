package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"strconv"
)

func AddRole(c *gin.Context) {
	var roleForm dto.RoleForm
	err := c.Bind(&roleForm)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	role, err := srv.AddRole(roleForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": role,
		"msg":  "ok",
	})
}

func GetRole(c *gin.Context) {
	Id := c.Param("id")
	roleId, _ := strconv.Atoi(Id)
	role, err := srv.GetRoleById(roleId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": role,
		"msg":  "ok",
	})
}

func EditRole(c *gin.Context) {
	var rf dto.RoleForm
	err := c.Bind(&rf)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	role, err := srv.GetRoleById(rf.RoleId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	role, err = srv.UpdateRole(role, rf)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": role,
		"msg":  "ok",
	})
}

func RoleList(c *gin.Context) {
	var roleSearch dto.SysRoleSearch
	err := c.Bind(&roleSearch)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	count, list, err := srv.GetSysRolePage(&roleSearch)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"list":      list,
			"count":     count,
			"pageIndex": roleSearch.PageIndex,
			"PageSize":  roleSearch.PageSize,
		},
		"msg": "ok",
	})
	return
}

func AllSysRoles(c *gin.Context) {
	list, err := srv.GetAllSysRoles()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": list,
		"msg":  "ok",
	})
}

func GetRoleMenus(c *gin.Context) {
	Id := c.Param("roleId")
	roleId, _ := strconv.Atoi(Id)
	if roleId < 0 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "角色ID错误",
		})
		return
	}

	sysMenus, err := srv.GetAllSysMenuLabels()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	roleMenuIds, err := srv.GetRoleSubMenuIds(roleId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"checkedKeys": roleMenuIds,
			"menus":       sysMenus,
		},
		"msg": "ok",
	})
}
