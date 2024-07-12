package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"memorandum/pkg/ctl"
	"memorandum/pkg/util"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// 检查请求头中是否包含 Token
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusNotFound,
				"msg":    "require parameter error",
				"data":   "lack of Token",
			})

			c.Abort()
			return
		}

		// 解析 Token
		claims, err := util.ParseToken(token)
		if err != nil {
			code, msg := handleTokenError(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": code,
				"msg":    "operate failed",
				"data":   msg,
			})
			c.Abort()
			return
		}

		// 将解析得到的用户信息添加到请求上下文中
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{ID: claims.Id}))
		c.Next()
	}
}

func handleTokenError(err error) (int, string) {
	var code int
	var msg string
	switch {
	case errors.Is(err, util.ErrTokenExpired):
		code = http.StatusUnauthorized
		msg = "token is expired"
	case errors.Is(err, util.ErrTokenInvalid):
		code = http.StatusUnauthorized
		msg = "token is invalid"
	default:
		code = http.StatusInternalServerError
		msg = "parse Token failed"
	}
	return code, msg
}
