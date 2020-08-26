package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/app/request"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"net/http"
)

// @Summary 添加/更新角色
// @Description 添加/更新角色
// @Tags 角色接口
// @Security ApiKeyAuth
// @ID AddRole
// @Accept application/json
// @Produce application/json
// @Param body body request.UpdateRoleForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/roles [post]
func UpdateRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.UpdateRoleForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	// 删除旧规则
	global.Enforcer.RemoveFilteredPolicy(0, form.RoleName)

	// 批量添加新规则
	for _, policy := range form.Policys {
		global.Enforcer.AddPolicy(form.RoleName, policy.Path, policy.Method)
	}

	appG.Response(http.StatusCreated, e.SUCCESS, nil)
}

// @Summary 为用户添加角色
// @Description 为用户添加角色
// @Tags 角色接口
// @Security ApiKeyAuth
// @ID AddRole
// @Accept application/json
// @Produce application/json
// @Param body body request.AddRoleForUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/user/role [post]
func AddRoleForUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AddRoleForUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	global.Enforcer.AddRoleForUser(form.UserName, form.RoleName)

	appG.Response(http.StatusCreated, e.SUCCESS, nil)
}

// @Summary 删除角色
// @Description 删除角色
// @Tags 角色接口
// @Security ApiKeyAuth
// @ID DeleteRole
// @Accept application/json
// @Produce application/json
// @Param role_name path string true "角色名"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/roles/{role_name} [delete]
func DeleteRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.DeleteRoleForm{RoleName: c.Params.ByName("role_name")}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	// 删除角色拥有的所有规则
	global.Enforcer.RemoveFilteredPolicy(0, form.RoleName)
	// 删除角色
	global.Enforcer.DeleteRole(form.RoleName)

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 查询所有角色
// @Description 查询所有角色
// @Tags 角色接口
// @Security ApiKeyAuth
// @ID GetAllRole
// @Accept application/json
// @Produce application/json
// @Success 200 {object} app.Response "success"
// @Router /api/v1/roles [get]
func GetAllRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"roles": global.Enforcer.GetAllRoles(),
	})
}
