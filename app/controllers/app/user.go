package app

import (
	"github.com/gin-gonic/gin"
	"study-gin/app/common/request"
	"study-gin/app/common/response"
	"study-gin/app/services"
)

// Register
// @Summary Register
// @Schemes
// @Description 用户注册
// @Tags example
// @Accept json
// @Produce json
// @param params body request.Register true "params"
// @Success 200 {object} response.Response
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
