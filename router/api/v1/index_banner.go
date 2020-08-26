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
	"github.com/zhenghuajing/fresh_shop/service/ad_service"
	"net/http"
)

// @Summary 添加首页轮播图
// @Description 添加首页轮播图
// @Tags 首页轮播图接口
// @Security ApiKeyAuth
// @ID AddIndexBanner
// @Accept application/json
// @Produce application/json
// @Param body body request.IndexBannerForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/index_banners [post]
func AddIndexBanner(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.IndexBannerForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	indexBannerModel := model.IndexBanner{
		ProductID: form.ProductID,
		Index:     form.Index,
		ImageUrl:  form.ImageUrl,
	}

	indexBanner, err := ad_service.AddIndexBanner(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_INDEX_BANNER_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, indexBanner)
}

// @Summary 删除首页轮播图
// @Description 删除首页轮播图
// @Tags 首页轮播图接口
// @Security ApiKeyAuth
// @ID DeleteIndexBanner
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/index_banners/{id} [delete]
func DeleteIndexBanner(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	indexBannerModel := model.IndexBanner{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := ad_service.ExistIndexBannerByID(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_INDEX_BANNER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_INDEX_BANNER, nil)
		return
	}

	if err := ad_service.DeleteIndexBanner(indexBannerModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_INDEX_BANNER_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新首页轮播图
// @Description 更新首页轮播图
// @Tags 首页轮播图接口
// @Security ApiKeyAuth
// @ID UpdateIndexBanner
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.IndexBannerForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/index_banners/{id} [put]
func UpdateIndexBanner(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.IndexBannerForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	indexBannerModel := model.IndexBanner{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		ProductID: form.ProductID,
		Index:     form.Index,
		ImageUrl:  form.ImageUrl,
	}

	ok, err := ad_service.ExistIndexBannerByID(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_INDEX_BANNER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_INDEX_BANNER, nil)
		return
	}

	indexBanner, err := ad_service.UpdateIndexBanner(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_INDEX_BANNER_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, indexBanner)
}

// @Summary 首页轮播图详情
// @Description 首页轮播图详情
// @Tags 首页轮播图接口
// @Security ApiKeyAuth
// @ID GetIndexBanner
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/index_banners/{id} [get]
func GetIndexBanner(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	indexBannerModel := model.IndexBanner{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := ad_service.ExistIndexBannerByID(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_INDEX_BANNER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_INDEX_BANNER, nil)
		return
	}

	indexBanner, err := ad_service.GetIndexBanner(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INDEX_BANNER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, indexBanner)
}

// @Summary 查询所有首页轮播图
// @Description 查询所有首页轮播图
// @Tags 首页轮播图接口
// @Security ApiKeyAuth
// @ID GetAllIndexBanner
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/index_banners [get]
func GetAllIndexBanner(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	indexBannerModel := model.IndexBanner{}

	index_banners, err := ad_service.GetIndexBanners(indexBannerModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_INDEX_BANNER_FAIL, nil)
		return
	}

	total, err := ad_service.GetIndexBannerTotal(indexBannerModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_INDEX_BANNER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": index_banners,
		"total": total,
	})
}
