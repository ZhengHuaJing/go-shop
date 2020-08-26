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
	"github.com/zhenghuajing/fresh_shop/service/trade_service"
	"net/http"
	"time"
)

// @Summary 添加订单信息
// @Description 添加订单信息
// @Tags 订单信息接口
// @Security ApiKeyAuth
// @ID AddOrderInfo
// @Accept application/json
// @Produce application/json
// @Param body body request.OrderInfoForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/orders [post]
func AddOrderInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.OrderInfoForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	nonceStr := util.MD5(fmt.Sprintf("%d%s", time.Now().Unix(), global.Config.MD5.Salt))
	// 创建订单
	orderInfoModel := model.OrderInfo{
		UserID:     form.UserID,
		OrderNo:    fmt.Sprintf("%d%d", time.Now().UnixNano(), form.UserID),
		NonceStr:   nonceStr,
		TradeNo:    "",
		PayStatus:  "待支付",
		PayType:    "支付宝",
		PostScript: form.PostScript,
		Money:      "",
		PayTime:    nil,
	}

	orderInfo, err := trade_service.AddOrderInfo(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ORDER_INFO_FAIL, nil)
		return
	}

	money := 0.00
	// 处理订单中每件商品
	for _, orderProduct := range form.OrderProducts {
		// 验证商品是否存在
		productModel := model.Product{
			Model: gorm.Model{
				ID: uint(orderProduct.ProductID),
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

		// 计算金额
		product, err := product_service.GetProduct(productModel)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_GET_PRODUCT_FAIL, nil)
			return
		}

		money += product.ShopPrice * float64(orderProduct.Num)

		orderProductModel := model.OrderProduct{
			OrderInfoID: int(orderInfo.ID),
			ProductID:   orderProduct.ProductID,
			Num:         orderProduct.Num,
		}
		if err = trade_service.AddOrderProduct(orderProductModel); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ORDER_PRODUCT_FAIL, nil)
			return
		}
	}
	orderInfo.Money = fmt.Sprintf("%.2f", money)

	// 更新订单信息
	orderInfo, err = trade_service.UpdateOrderInfo(*orderInfo)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ORDER_INFO_FAIL, nil)
		return
	}

	alipayUrl := ""
	// 调用支付接口
	switch form.ClientType {
	case 1:
		alipayUrl = util.WebPageAlipay(form.ReturnUrl, orderInfo.OrderNo, orderInfo.NonceStr, orderInfo.Money)
	case 2:
		alipayUrl = util.WapAlipay(form.ReturnUrl, orderInfo.OrderNo, orderInfo.NonceStr, orderInfo.Money)
	}

	appG.Response(http.StatusCreated, e.SUCCESS, map[string]string{
		"alipay_url": alipayUrl,
	})
}

// @Summary 删除订单信息
// @Description 删除订单信息
// @Tags 订单信息接口
// @Security ApiKeyAuth
// @ID DeleteOrderInfo
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/orders/{id} [delete]
func DeleteOrderInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	orderInfoModel := model.OrderInfo{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := trade_service.ExistOrderInfoByID(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ORDER_INFO_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_ORDER_INFO, nil)
		return
	}

	if err := trade_service.DeleteOrderInfo(orderInfoModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ORDER_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新订单信息
// @Description 更新订单信息
// @Tags 订单信息接口
// @Security ApiKeyAuth
// @ID UpdateOrderInfo
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.OrderInfoForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/orders/{id} [put]
func UpdateOrderInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.OrderInfoForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	orderInfoModel := model.OrderInfo{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := trade_service.ExistOrderInfoByID(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ORDER_INFO_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_ORDER_INFO, nil)
		return
	}

	orderInfo, err := trade_service.UpdateOrderInfo(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ORDER_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, orderInfo)
}

// @Summary 订单信息详情
// @Description 订单信息详情
// @Tags 订单信息接口
// @Security ApiKeyAuth
// @ID GetOrderInfo
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/orders/{id} [get]
func GetOrderInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	orderInfoModel := model.OrderInfo{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := trade_service.ExistOrderInfoByID(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ORDER_INFO_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_ORDER_INFO, nil)
		return
	}

	orderInfo, err := trade_service.GetOrderInfo(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ORDER_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, orderInfo)
}

// @Summary 查询所有订单信息
// @Description 查询所有订单信息
// @Tags 订单信息接口
// @Security ApiKeyAuth
// @ID GetAllOrderInfo
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/orders [get]
func GetAllOrderInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	orderInfoModel := model.OrderInfo{}

	orderInfos, err := trade_service.GetOrderInfos(orderInfoModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_ORDER_INFO_FAIL, nil)
		return
	}

	total, err := trade_service.GetOrderInfoTotal(orderInfoModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ORDER_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": orderInfos,
		"total": total,
	})
}

// @Summary 支付成功回调接口
// @Description 支付成功回调接口
// @Tags 订单信息接口
// @ID AddOrderInfo
// @Accept application/json
// @Produce application/json
// @Success 200 {object} app.Response "success"
// @Router /api/v1/pay [post]
func Pay(c *gin.Context) {
	tradeStatus := c.DefaultPostForm("trade_status", "")
	outTradeNo := c.DefaultPostForm("out_trade_no", "")
	tradeNo := c.DefaultPostForm("trade_no", "")
	nonceStr := c.DefaultPostForm("passback_params", "")

	orderInfoModel := model.OrderInfo{
		OrderNo: outTradeNo,
	}

	orderInfo, err := trade_service.GetOrderInfoByOrderNo(orderInfoModel)
	if err != nil {
		global.Log.Error(err)
	}

	if orderInfo == nil {
		return
	}

	if tradeStatus == "TRADE_SUCCESS" && orderInfo.NonceStr == nonceStr {
		orderInfo.PayStatus = "已支付"
		orderInfo.TradeNo = tradeNo
		// 更新订单状态
		trade_service.UpdateOrderInfo(*orderInfo)
	}
}
