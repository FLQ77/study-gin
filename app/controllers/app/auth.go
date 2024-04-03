package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"study-gin/app/common/request"
	"study-gin/app/common/response"
	"study-gin/app/services"
)

// @Summary Login
// @Description 用户登录
// @Tags example
// @Accept json
// @Produce json
// @Param params body request.Login true "params"
// @Success 200 {object} response.Response "{"error_code": 0, "message": "OK", "data": {}}"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// @Summary Info
// @Description 获取用户信息
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "{"error_code": 0, "message": "OK", "data": {}}"
// @Router /auth/info [get]
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

// @Summary Logout
// @Description 用户登出
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "{"error_code": 0, "message": "OK", "data": {}}"
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}
