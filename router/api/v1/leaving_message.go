package v1

import (
	"github.com/gin-gonic/gin"
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

// @Summary 添加用户留言
// @Description 添加用户留言
// @Tags 用户留言接口
// @Security ApiKeyAuth
// @ID AddLeavingMessage
// @Accept application/json
// @Produce application/json
// @Param body body request.LeavingMessageForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/leaving_messages [post]
func AddLeavingMessage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.LeavingMessageForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	leavingMessageModel := model.LeavingMessage{
		UserID:      form.UserID,
		MessageType: form.MessageType,
		Title:       form.Title,
		Content:     form.Content,
	}

	leavingMessage, err := user_service.AddLeavingMessage(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_LEAVING_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, leavingMessage)
}

// @Summary 删除用户留言
// @Description 删除用户留言
// @Tags 用户留言接口
// @Security ApiKeyAuth
// @ID DeleteLeavingMessage
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/leaving_messages/{id} [delete]
func DeleteLeavingMessage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	leavingMessageModel := model.LeavingMessage{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistLeavingMessageByID(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_LEAVING_MESSAGE_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_LEAVING_MESSAGE, nil)
		return
	}

	if err := user_service.DeleteLeavingMessage(leavingMessageModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_LEAVING_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新用户留言
// @Description 更新用户留言
// @Tags 用户留言接口
// @Security ApiKeyAuth
// @ID UpdateLeavingMessage
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.LeavingMessageForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/leaving_messages/{id} [put]
func UpdateLeavingMessage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.LeavingMessageForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	leavingMessageModel := model.LeavingMessage{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		UserID:      form.UserID,
		MessageType: form.MessageType,
		Title:       form.Title,
		Content:     form.Content,
	}

	ok, err := user_service.ExistLeavingMessageByID(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_LEAVING_MESSAGE_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_LEAVING_MESSAGE, nil)
		return
	}

	leavingMessage, err := user_service.UpdateLeavingMessage(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_LEAVING_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, leavingMessage)
}

// @Summary 用户留言详情
// @Description 用户留言详情
// @Tags 用户留言接口
// @Security ApiKeyAuth
// @ID GetLeavingMessage
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/leaving_messages/{id} [get]
func GetLeavingMessage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	leavingMessageModel := model.LeavingMessage{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistLeavingMessageByID(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_LEAVING_MESSAGE_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_LEAVING_MESSAGE, nil)
		return
	}

	leavingMessage, err := user_service.GetLeavingMessage(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_LEAVING_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, leavingMessage)
}

// @Summary 查询所有用户留言
// @Description 查询所有用户留言
// @Tags 用户留言接口
// @Security ApiKeyAuth
// @ID GetAllLeavingMessage
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/leaving_messages [get]
func GetAllLeavingMessage(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	leavingMessageModel := model.LeavingMessage{}

	leavingMessages, err := user_service.GetLeavingMessages(leavingMessageModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_LEAVING_MESSAGE_FAIL, nil)
		return
	}

	total, err := user_service.GetLeavingMessageTotal(leavingMessageModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_LEAVING_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": leavingMessages,
		"total": total,
	})
}
