package common

import (
	"github.com/gin-gonic/gin"
	"study-gin/app/common/request"
	"study-gin/app/common/response"
	"study-gin/app/services"
)

// ImageUpload 图片上传
// @Tags 公共
// @Summary 图片上传
// @Produce  json
// @Param params body request.ImageUpload true "params"
// @Success 200 {object} response.Response
// @Router /api/image_upload [post]
func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
