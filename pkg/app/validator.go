package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int, map[string]string) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS, nil
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR, nil
	}
	if !check {
		markErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS, formatErrors(valid.Errors)
	}

	return http.StatusOK, e.SUCCESS, nil
}

// MarkErrors logs error logs
func markErrors(errors []*validation.Error) {
	for _, err := range errors {
		global.Log.Error(err.Key, err.Message)
	}

	return
}

// 格式化错误信息，作为response返回
func formatErrors(errors []*validation.Error) map[string]string {
	errMsg := map[string]string{}

	for _, err := range errors {
		errMsg[err.Key] = err.Message[1:] // 去掉开头第一个空格
	}

	return errMsg
}
