package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/zhenghuajing/fresh_shop/model"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/app/request"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	"github.com/zhenghuajing/fresh_shop/service/user_service"
	"math/rand"
	"net/http"
)

// @Summary 生成邮箱验证码
// @Description 生成邮箱验证码
// @Tags 验证码接口
// @Security ApiKeyAuth
// @ID AddVerifyCode
// @Accept application/json
// @Produce application/json
// @Param body body request.VerifyCodeForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/verify_code [post]
func GenerateVerifyCode(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.VerifyCodeForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	// 发送邮件
	code := rand.Intn(900000) + 100000
	msg := fmt.Sprintf("您的验证码为：<b>%d</b>", code)
	go util.SendEmail(form.Email, msg)

	verifyCodeModel := model.VerifyCode{
		Email: form.Email,
		Code:  com.ToStr(code),
	}

	if err := user_service.AddVerifyCode(verifyCodeModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_VERIFY_CODE_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, nil)
}
