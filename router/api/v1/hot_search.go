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
	"github.com/zhenghuajing/fresh_shop/service/hot_search_service"
	"net/http"
)

// @Summary 添加热搜词
// @Description 添加热搜词
// @Tags 热搜词接口
// @Security ApiKeyAuth
// @ID AddHotSearch
// @Accept application/json
// @Produce application/json
// @Param body body request.HotSearchForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/hot_searchs [post]
func AddHotSearch(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.HotSearchForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	hotSearchModel := model.HotSearch{
		Keyword: form.Keyword,
		Index:   form.Index,
	}

	hotSearch, err := hot_search_service.AddHotSearch(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_HOT_SEARCH_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, hotSearch)
}

// @Summary 删除热搜词
// @Description 删除热搜词
// @Tags 热搜词接口
// @Security ApiKeyAuth
// @ID DeleteHotSearch
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/hot_searchs/{id} [delete]
func DeleteHotSearch(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	hotSearchModel := model.HotSearch{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := hot_search_service.ExistHotSearchByID(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOT_SEARCH_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_HOT_SEARCH, nil)
		return
	}

	if err := hot_search_service.DeleteHotSearch(hotSearchModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_HOT_SEARCH_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新热搜词
// @Description 更新热搜词
// @Tags 热搜词接口
// @Security ApiKeyAuth
// @ID UpdateHotSearch
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.HotSearchForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/hot_searchs/{id} [put]
func UpdateHotSearch(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.HotSearchForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	hotSearchModel := model.HotSearch{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		Keyword: form.Keyword,
		Index:   form.Index,
	}

	ok, err := hot_search_service.ExistHotSearchByID(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOT_SEARCH_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_HOT_SEARCH, nil)
		return
	}

	hotSearch, err := hot_search_service.UpdateHotSearch(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_HOT_SEARCH_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, hotSearch)
}

// @Summary 热搜词详情
// @Description 热搜词详情
// @Tags 热搜词接口
// @Security ApiKeyAuth
// @ID GetHotSearch
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/hot_searchs/{id} [get]
func GetHotSearch(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	hotSearchModel := model.HotSearch{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := hot_search_service.ExistHotSearchByID(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOT_SEARCH_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_HOT_SEARCH, nil)
		return
	}

	hotSearch, err := hot_search_service.GetHotSearch(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_HOT_SEARCH_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, hotSearch)
}

// @Summary 查询所有热搜词
// @Description 查询所有热搜词
// @Tags 热搜词接口
// @Security ApiKeyAuth
// @ID GetAllHotSearch
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/hot_searchs [get]
func GetAllHotSearch(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	hotSearchModel := model.HotSearch{}

	hotSearchs, err := hot_search_service.GetHotSearchs(hotSearchModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_HOT_SEARCH_FAIL, nil)
		return
	}

	total, err := hot_search_service.GetHotSearchTotal(hotSearchModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_HOT_SEARCH_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": hotSearchs,
		"total": total,
	})
}
