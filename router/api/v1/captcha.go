package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

// @Summary 生成验证码
// @Description 生成验证码
// @Tags 验证码接口
// @ID Captcha
// @Accept application/json
// @Produce application/json
// @Success 200 {object} app.Response "success"
// @Router /api/v1/captcha [get]
func GenerateCaptcha(c *gin.Context) {
	appG := app.Gin{C: c}
	captchaCfg := global.Config.Captcha
	driver := base64Captcha.NewDriverDigit(captchaCfg.ImgHeight, captchaCfg.ImgWidth, captchaCfg.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CAPTCHA_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"captcha_id":   id,
		"image_base64": b64s,
	})
	return
}
