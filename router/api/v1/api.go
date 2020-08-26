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
	"github.com/zhenghuajing/fresh_shop/service/api_service"
	"net/http"
)

//@Summary API详情
//@Description API详情
//@Tags API接口
//@Security ApiKeyAuth
//@ID GetApi
//@Accept application/json
//@Produce application/json
//@Param id path int true "id"
//@Success 200 {object} app.Response "success"
//@Router /api/v1/apis/{id} [get]
func GetApi(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.ID{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	apiModel := model.Api{
		Model: gorm.Model{
			ID: uint(form.ID),
		},
	}

	ok, err := api_service.ExistApiByID(apiModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_API_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_API, nil)
		return
	}

	api, err := api_service.GetApi(apiModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_API_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, api)
}

// @Summary 查询所有API
// @Description 查询所有API
// @Tags API接口
// @Security ApiKeyAuth
// @ID GetAllApi
// @Accept application/json
// @Produce application/json
// @Param page query int false "page"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/apis [get]
func GetAllApi(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := global.Config.App.PageSize

	apiModel := model.Api{}

	apis, err := api_service.GetApis(apiModel, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ALL_API_FAIL, nil)
		return
	}

	total, err := api_service.GetApiTotal(apiModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_API_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": apis,
		"total": total,
	})
}
