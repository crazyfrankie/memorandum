package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"memorandum/repository/db/dao"
	"memorandum/repository/db/model"
	"memorandum/service"
)

var userService service.UserService = service.NewUserService(dao.NewUserRepository())

// RegisterHandler 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body model.LoginData true "注册信息"
// @Success 200 {string} string "register successfully"
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/user/register [post]
func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.LoginData
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		if err := userService.RegisterUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, "register successfully")
	}
}

// LoginHandler 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body model.LoginData true "登录信息"
// @Success 200 {object} ctl.TokenResponse
// @Failure 400 {object} ctl.ErrResponse
// @Failure 401 {object} ctl.ErrResponse
// @Router /v1/user/login [post]
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData model.LoginData
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		user, token, err := userService.LoginUser(&loginData)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, TokenResponse("Login successful", user, token))

	}
}
