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
	"github.com/zhenghuajing/fresh_shop/service/trade_service"
	"net/http"
)

// @Summary 添加购物车
// @Description 添加购物车
// @Tags 购物车接口
// @Security ApiKeyAuth
// @ID AddShoppingCart
// @Accept application/json
// @Produce application/json
// @Param body body request.ShoppingCartForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/shopping_carts [post]
func AddShoppingCart(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.ShoppingCartForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	shoppingCartModel := model.ShoppingCart{
		UserID:    form.UserID,
		ProductID: form.ProductID,
		Num:       form.Num,
	}

	shoppingCart, err := trade_service.AddShoppingCart(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_SHOPPING_CART_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, shoppingCart)
}

// @Summary 删除购物车
// @Description 删除购物车
// @Tags 购物车接口
// @Security ApiKeyAuth
// @ID DeleteShoppingCart
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/shopping_carts/{id} [delete]
func DeleteShoppingCart(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	shoppingCartModel := model.ShoppingCart{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := trade_service.ExistShoppingCartByID(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_SHOPPING_CART_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_SHOPPING_CART, nil)
		return
	}

	if err := trade_service.DeleteShoppingCart(shoppingCartModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_SHOPPING_CART_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新购物车
// @Description 更新购物车
// @Tags 购物车接口
// @Security ApiKeyAuth
// @ID UpdateShoppingCart
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.ShoppingCartForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/shopping_carts/{id} [put]
func UpdateShoppingCart(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ShoppingCartForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	shoppingCartModel := model.ShoppingCart{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		UserID:    form.UserID,
		ProductID: form.ProductID,
		Num:       form.Num,
	}

	ok, err := trade_service.ExistShoppingCartByID(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_SHOPPING_CART_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_SHOPPING_CART, nil)
		return
	}

	shoppingCart, err := trade_service.UpdateShoppingCart(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_SHOPPING_CART_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, shoppingCart)
}

// @Summary 购物车详情
// @Description 购物车详情
// @Tags 购物车接口
// @Security ApiKeyAuth
// @ID GetShoppingCart
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/shopping_carts/{id} [get]
func GetShoppingCart(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	shoppingCartModel := model.ShoppingCart{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := trade_service.ExistShoppingCartByID(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_SHOPPING_CART_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_SHOPPING_CART, nil)
		return
	}

	shoppingCart, err := trade_service.GetShoppingCart(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SHOPPING_CART_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, shoppingCart)
}

// @Summary 查询所有购物车
// @Description 查询所有购物车
// @Tags 购物车接口
// @Security ApiKeyAuth
// @ID GetAllShoppingCart
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/shopping_carts [get]
func GetAllShoppingCart(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	shoppingCartModel := model.ShoppingCart{}

	shoppingCarts, err := trade_service.GetShoppingCarts(shoppingCartModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_SHOPPING_CART_FAIL, nil)
		return
	}

	total, err := trade_service.GetShoppingCartTotal(shoppingCartModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_SHOPPING_CART_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": shoppingCarts,
		"total": total,
	})
}
