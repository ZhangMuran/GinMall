package middleware

import (
	"fmt"
	"time"

	"github.com/ZhangMuran/GinMall/service/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := utils.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			fmt.Println("here token =", token)
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = utils.ErrAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = utils.ErrTokenTimeOut
			}
		}
		if code != utils.Success {
			c.JSON(200, utils.Response{
				Errno: code,
				ErrMsg: utils.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}