package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"github.com/zhenghuajing/fresh_shop/pkg/upload"
	"net/http"
)

// @Summary 上传图片
// @Description 上传图片
// @Tags 上传接口
// @ID UploadImage
// @Accept multipart/form-data
// @Produce application/json
// @Param image formData file true "图片"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/upload [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	if err := upload.CheckImage(fullPath); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
