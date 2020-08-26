package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/app/request"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	"github.com/zhenghuajing/fresh_shop/service/user_service"
	"net/http"
)

// @Summary 添加用户
// @Description 添加用户
// @Tags 用户接口
// @ID AddUser
// @Accept application/json
// @Produce application/json
// @Param body body request.AddUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users [post]
func AddUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AddUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	// 校验验证码
	verifyCodeModel := model.VerifyCode{
		Code:  form.VerifyCode,
		Email: form.UserName,
	}
	ok, err := user_service.CheckVerifyCode(verifyCodeModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_VERIFY_CODE_FAIL, nil)
		return
	}

	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_VERIFY_CODE, nil)
		return
	}

	newUUID, _ := uuid.NewUUID()

	userModel := model.User{
		UserName: form.UserName,
		Password: util.MD5(form.Password + global.Config.MD5.Salt),
		UUID:     newUUID,
		IP:       util.RemoteIP(c.Request),
	}

	ok, err = user_service.ExistUserByUserName(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if ok {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_USER_NAME, nil)
		return
	}

	user, err := user_service.AddUser(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	// 默认角色为普通用户
	global.Enforcer.AddRoleForUser(user.UserName, global.Config.Casbin.RoleCommonUser)

	appG.Response(http.StatusCreated, e.SUCCESS, user)
}

// @Summary 删除用户
// @Description 删除用户
// @Tags 用户接口
// @Security ApiKeyAuth
// @ID DeleteUser
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	userModel := model.User{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistUserByID(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	if err := user_service.DeleteUser(userModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新用户
// @Description 更新用户
// @Tags 用户接口
// @Security ApiKeyAuth
// @ID UpdateUser
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.UpdateUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.UpdateUserForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	userModel := model.User{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		Gender:   form.Gender,
		RealName: form.RealName,
		Mobile:   form.Mobile,
	}

	if form.Password != "" {
		if len(form.Password) < 6 {
			appG.Response(http.StatusBadRequest, e.ERROR_PASSWORD_SIZE, nil)
			return
		}

		userModel.Password = util.MD5(form.Password + global.Config.MD5.Salt)
	}

	ok, err := user_service.ExistUserByID(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user, err := user_service.UpdateUser(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, user)
}

// @Summary 用户详情
// @Description 用户详情
// @Tags 用户接口
// @Security ApiKeyAuth
// @ID GetUser
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	userModel := model.User{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistUserByID(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user, err := user_service.GetUser(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// @Summary 查询所有用户
// @Description 查询所有用户
// @Tags 用户接口
// @Security ApiKeyAuth
// @ID GetAllUser
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users [get]
func GetAllUser(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	userModel := model.User{}

	users, err := user_service.GetUsers(userModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_USER_FAIL, nil)
		return
	}

	total, err := user_service.GetUserTotal(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": users,
		"total": total,
	})
}

// @Summary 用户登录
// @Description 用户登录
// @Tags 用户接口
// @ID UserLogin
// @Accept application/json
// @Produce application/json
// @Param body body request.LoginUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.LoginUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	userModel := model.User{
		UserName: form.UserName,
		Password: util.MD5(form.Password + global.Config.MD5.Salt),
		IP:       util.RemoteIP(c.Request),
	}

	user, err := user_service.UserLogin(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}

	if user == nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_NAME_OR_PASSWORD, nil)
		return
	}

	token, err := util.GenerateToken(int(user.ID), user.UserName, user.Password, user.UUID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_TOKEN_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
