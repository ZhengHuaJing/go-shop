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

// @Summary 添加用户收货地址
// @Description 添加用户收货地址
// @Tags 用户收货地址接口
// @Security ApiKeyAuth
// @ID AddAddress
// @Accept application/json
// @Produce application/json
// @Param body body request.AddressForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/addresses [post]
func AddAddress(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AddressForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	addressModel := model.Address{
		UserID:        form.UserID,
		Province:      form.Province,
		City:          form.City,
		District:      form.District,
		DetailAddress: form.DetailAddress,
		SignerName:    form.SignerName,
		SignerMobile:  form.SignerMobile,
	}

	address, err := user_service.AddAddress(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ADDRESS_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, address)
}

// @Summary 删除用户收货地址
// @Description 删除用户收货地址
// @Tags 用户收货地址接口
// @Security ApiKeyAuth
// @ID DeleteAddress
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/addresses/{id} [delete]
func DeleteAddress(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	addressModel := model.Address{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistAddressByID(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ADDRESS_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_ADDRESS, nil)
		return
	}

	if err := user_service.DeleteAddress(addressModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ADDRESS_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新用户收货地址
// @Description 更新用户收货地址
// @Tags 用户收货地址接口
// @Security ApiKeyAuth
// @ID UpdateAddress
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.AddressForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/addresses/{id} [put]
func UpdateAddress(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.AddressForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	addressModel := model.Address{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		UserID:        form.UserID,
		Province:      form.Province,
		City:          form.City,
		District:      form.District,
		DetailAddress: form.DetailAddress,
		SignerName:    form.SignerName,
		SignerMobile:  form.SignerMobile,
	}

	ok, err := user_service.ExistAddressByID(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ADDRESS_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_ADDRESS, nil)
		return
	}

	address, err := user_service.UpdateAddress(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ADDRESS_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, address)
}

// @Summary 用户收货地址详情
// @Description 用户收货地址详情
// @Tags 用户收货地址接口
// @Security ApiKeyAuth
// @ID GetAddress
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/addresses/{id} [get]
func GetAddress(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	addressModel := model.Address{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := user_service.ExistAddressByID(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ADDRESS_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_ADDRESS, nil)
		return
	}

	address, err := user_service.GetAddress(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ADDRESS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, address)
}

// @Summary 查询所有用户收货地址
// @Description 查询所有用户收货地址
// @Tags 用户收货地址接口
// @Security ApiKeyAuth
// @ID GetAllAddress
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/addresses [get]
func GetAllAddress(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	addressModel := model.Address{}

	addresss, err := user_service.GetAddresses(addressModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_ADDRESS_FAIL, nil)
		return
	}

	total, err := user_service.GetAddressTotal(addressModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ADDRESS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": addresss,
		"total": total,
	})
}
