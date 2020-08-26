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

// @Summary 添加收藏
// @Description 添加收藏
// @Tags 用户收藏接口
// @Security ApiKeyAuth
// @ID AddCollect
// @Accept application/json
// @Produce application/json
// @Param body body request.CollectForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/collects [post]
func AddCollect(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.CollectForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	collectModel := model.Collect{
		UserID:    form.UserID,
		ProductID: form.ProductID,
	}

	collect, err := user_service.AddCollect(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_COLLECT_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, collect)
}

// @Summary 删除收藏
// @Description 删除收藏
// @Tags 用户收藏接口
// @Security ApiKeyAuth
// @ID DeleteCollect
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/collects/{id} [delete]
func DeleteCollect(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	collectModel := model.Collect{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistCollectByID(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COLLECT_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_COLLECT, nil)
		return
	}

	if err := user_service.DeleteCollect(collectModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_COLLECT_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新收藏
// @Description 更新收藏
// @Tags 用户收藏接口
// @Security ApiKeyAuth
// @ID UpdateCollect
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.CollectForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/collects/{id} [put]
func UpdateCollect(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.CollectForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	collectModel := model.Collect{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		UserID:    form.UserID,
		ProductID: form.ProductID,
	}

	ok, err := user_service.ExistCollectByID(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COLLECT_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_COLLECT, nil)
		return
	}

	collect, err := user_service.UpdateCollect(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_COLLECT_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, collect)
}

// @Summary 收藏详情
// @Description 收藏详情
// @Tags 用户收藏接口
// @Security ApiKeyAuth
// @ID GetCollect
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/collects/{id} [get]
func GetCollect(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	collectModel := model.Collect{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistCollectByID(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COLLECT_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_COLLECT, nil)
		return
	}

	collect, err := user_service.GetCollect(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_COLLECT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, collect)
}

// @Summary 查询所有收藏
// @Description 查询所有收藏
// @Tags 用户收藏接口
// @Security ApiKeyAuth
// @ID GetAllCollect
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/collects [get]
func GetAllCollect(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	collectModel := model.Collect{}

	collects, err := user_service.GetCollects(collectModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_COLLECT_FAIL, nil)
		return
	}

	total, err := user_service.GetCollectTotal(collectModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_COLLECT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": collects,
		"total": total,
	})
}
