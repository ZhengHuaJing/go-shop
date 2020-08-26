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
	"github.com/zhenghuajing/fresh_shop/service/product_service"
	"net/http"
)

// @Summary 添加分类
// @Description 添加分类
// @Tags 分类接口
// @Security ApiKeyAuth
// @ID AddCategory
// @Accept application/json
// @Produce application/json
// @Param body body request.CategoryForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/categories [post]
func AddCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.CategoryForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	categoryModel := model.Category{
		ParentCategoryID: form.ParentCategoryID,
		Name:             form.Name,
		Code:             form.Code,
		Desc:             form.Desc,
		CategoryType:     form.CategoryType,
		IsTab:            form.IsTab,
	}

	category, err := product_service.AddCategory(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, category)
}

// @Summary 删除分类
// @Description 删除分类
// @Tags 分类接口
// @Security ApiKeyAuth
// @ID DeleteCategory
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	categoryModel := model.Category{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := product_service.ExistCategoryByID(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	if err := product_service.DeleteCategory(categoryModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}

// @Summary 更新分类
// @Description 更新分类
// @Tags 分类接口
// @Security ApiKeyAuth
// @ID UpdateCategory
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param body body request.CategoryForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.CategoryForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	categoryModel := model.Category{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
		ParentCategoryID: form.ParentCategoryID,
		Name:             form.Name,
		Code:             form.Code,
		Desc:             form.Desc,
		CategoryType:     form.CategoryType,
		IsTab:            form.IsTab,
	}

	ok, err := product_service.ExistCategoryByID(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	category, err := product_service.UpdateCategory(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, category)
}

// @Summary 分类详情
// @Description 分类详情
// @Tags 分类接口
// @Security ApiKeyAuth
// @ID GetCategory
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/categories/{id} [get]
func GetCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	categoryModel := model.Category{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := product_service.ExistCategoryByID(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	category, err := product_service.GetCategory(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, category)
}

// @Summary 查询所有分类
// @Description 查询所有分类
// @Tags 分类接口
// @Security ApiKeyAuth
// @ID GetAllCategory
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/categories [get]
func GetAllCategory(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	categoryModel := model.Category{}

	categorys, err := product_service.GetCategorys(categoryModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_CATEGORY_FAIL, nil)
		return
	}

	total, err := product_service.GetCategoryTotal(categoryModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": categorys,
		"total": total,
	})
}
