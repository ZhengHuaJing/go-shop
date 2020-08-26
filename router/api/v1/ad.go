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

// @Summary 添加分类广告商品
// @Description 添加分类广告商品
// @Tags 分类广告商品接口
// @Security ApiKeyAuth
// @ID AddAd
// @Accept application/json
// @Produce application/json
// @Param body body request.AdForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/ads [post]
func AddAd(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AdForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	adModel := model.Ad{
		CategoryID: form.CategoryID,
		ProductID:  form.ProductID,
	}

	ad, err := ad_service.AddAd(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_AD_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, ad)
}

// @Summary 删除分类广告商品
// @Description 删除分类广告商品
// @Tags 分类广告商品接口
// @Security ApiKeyAuth
// @ID DeleteAd
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/ads/{id} [delete]
func DeleteAd(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	adModel := model.Ad{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := ad_service.ExistAdByID(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_AD_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_AD, nil)
		return
	}

	if err := ad_service.DeleteAd(adModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_AD_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新分类广告商品
// @Description 更新分类广告商品
// @Tags 分类广告商品接口
// @Security ApiKeyAuth
// @ID UpdateAd
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.AdForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/ads/{id} [put]
func UpdateAd(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.AdForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	adModel := model.Ad{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		CategoryID: form.CategoryID,
		ProductID:  form.ProductID,
	}

	ok, err := ad_service.ExistAdByID(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_AD_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_AD, nil)
		return
	}

	ad, err := ad_service.UpdateAd(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_AD_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, ad)
}

// @Summary 分类广告商品详情
// @Description 分类广告商品详情
// @Tags 分类广告商品接口
// @Security ApiKeyAuth
// @ID GetAd
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/ads/{id} [get]
func GetAd(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	adModel := model.Ad{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := ad_service.ExistAdByID(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_AD_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_AD, nil)
		return
	}

	ad, err := ad_service.GetAd(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, ad)
}

// @Summary 查询所有分类广告商品
// @Description 查询所有分类广告商品
// @Tags 分类广告商品接口
// @Security ApiKeyAuth
// @ID GetAllAd
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/ads [get]
func GetAllAd(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	adModel := model.Ad{}

	ads, err := ad_service.GetAds(adModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_AD_FAIL, nil)
		return
	}

	total, err := ad_service.GetAdTotal(adModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_AD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": ads,
		"total": total,
	})
}
