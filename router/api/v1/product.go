package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/app/request"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	"github.com/zhenghuajing/fresh_shop/service/product_service"
	"net/http"
	"strings"
	"time"
)

// @Summary 添加商品
// @Description 添加商品
// @Tags 商品接口
// @Security ApiKeyAuth
// @ID AddProduct
// @Accept application/json
// @Produce application/json
// @Param body body request.ProductForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/products [post]
func AddProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.ProductForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	productModel := model.Product{
		CategoryID:    form.CategoryID,
		FrontImageUrl: form.FrontImageUrl,
		ProductSN:     fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:          form.Name,
		ClickNum:      0,
		SoldNum:       0,
		CollectNum:    0,
		LeftNum:       form.LeftNum,
		MarketPrice:   form.MarketPrice,
		ShopPrice:     form.ShopPrice,
		Brief:         form.Brief,
		Desc:          form.Desc,
		IsFreight:     form.IsFreight,
		IsNew:         form.IsNew,
		IsHot:         form.IsHot,
	}
	productModel.BannerUrlStrs = strings.Join(form.BannerUrls, ";")

	product, err := product_service.AddProduct(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, product)
}

// @Summary 删除商品
// @Description 删除商品
// @Tags 商品接口
// @Security ApiKeyAuth
// @ID DeleteProduct
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	productModel := model.Product{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := product_service.ExistProductByID(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	if err := product_service.DeleteProduct(productModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新商品
// @Description 更新商品
// @Tags 商品接口
// @Security ApiKeyAuth
// @ID UpdateProduct
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.ProductForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ProductForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	productModel := model.Product{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		CategoryID: form.CategoryID,
		//BannerUrls:    form.BannerUrls,
		FrontImageUrl: form.FrontImageUrl,
		Name:          form.Name,
		LeftNum:       form.LeftNum,
		MarketPrice:   form.MarketPrice,
		ShopPrice:     form.ShopPrice,
		Brief:         form.Brief,
		Desc:          form.Desc,
		IsFreight:     form.IsFreight,
		IsNew:         form.IsNew,
		IsHot:         form.IsHot,
	}
	productModel.BannerUrlStrs = strings.Join(productModel.BannerUrls, ";")

	ok, err := product_service.ExistProductByID(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	product, err := product_service.UpdateProduct(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, product)
}

// @Summary 商品详情
// @Description 商品详情
// @Tags 商品接口
// @ID GetProduct
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/products/{id} [get]
func GetProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	productModel := model.Product{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := product_service.ExistProductByID(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	product, err := product_service.GetProduct(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, product)
}

// @Summary 查询所有商品
// @Description 查询所有商品
// @Tags 商品接口
// @ID GetAllProduct
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/products [get]
func GetAllProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	productModel := model.Product{}

	products, err := product_service.GetProducts(productModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_PRODUCT_FAIL, nil)
		return
	}

	total, err := product_service.GetProductTotal(productModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": products,
		"total": total,
	})
}
