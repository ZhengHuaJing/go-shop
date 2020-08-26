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

// @Summary 添加品牌
// @Description 添加品牌
// @Tags 品牌接口
// @Security ApiKeyAuth
// @ID AddBrand
// @Accept application/json
// @Produce application/json
// @Param body body request.BrandForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/brands [post]
func AddBrand(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.BrandForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
	}

	brandModel := model.Brand{
		CategoryID: form.CategoryID,
		Name:       form.Name,
		Desc:       form.Desc,
		ImageUrl:   form.ImageUrl,
	}

	brand, err := ad_service.AddBrand(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_BRAND_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, brand)
}

// @Summary 删除品牌
// @Description 删除品牌
// @Tags 品牌接口
// @Security ApiKeyAuth
// @ID DeleteBrand
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/brands/{id} [delete]
func DeleteBrand(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	brandModel := model.Brand{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := ad_service.ExistBrandByID(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_BRAND_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_BRAND, nil)
		return
	}

	if err := ad_service.DeleteBrand(brandModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_BRAND_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新品牌
// @Description 更新品牌
// @Tags 品牌接口
// @Security ApiKeyAuth
// @ID UpdateBrand
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.BrandForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/brands/{id} [put]
func UpdateBrand(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.BrandForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	brandModel := model.Brand{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		CategoryID: form.CategoryID,
		Name:       form.Name,
		Desc:       form.Desc,
		ImageUrl:   form.ImageUrl,
	}

	ok, err := ad_service.ExistBrandByID(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_BRAND_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_BRAND, nil)
		return
	}

	brand, err := ad_service.UpdateBrand(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_BRAND_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, brand)
}

// @Summary 品牌详情
// @Description 品牌详情
// @Tags 品牌接口
// @Security ApiKeyAuth
// @ID GetBrand
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/brands/{id} [get]
func GetBrand(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	brandModel := model.Brand{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := ad_service.ExistBrandByID(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_BRAND_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_BRAND, nil)
		return
	}

	brand, err := ad_service.GetBrand(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_BRAND_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, brand)
}

// @Summary 查询所有品牌
// @Description 查询所有品牌
// @Tags 品牌接口
// @Security ApiKeyAuth
// @ID GetAllBrand
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/brands [get]
func GetAllBrand(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	brandModel := model.Brand{}

	brands, err := ad_service.GetBrands(brandModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_BRAND_FAIL, nil)
		return
	}

	total, err := ad_service.GetBrandTotal(brandModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_BRAND_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": brands,
		"total": total,
	})
}
