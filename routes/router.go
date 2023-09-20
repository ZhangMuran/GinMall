package routes

import (
	"net/http"

	"github.com/ZhangMuran/GinMall/controller/user"
	"github.com/ZhangMuran/GinMall/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	userGroup := r.Group("/user")
	{
		userGroup.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		
		// 用户注册
		userGroup.POST("register", user.UserRegister)
		//用户登录
		userGroup.POST("login", user.UserLogin)

		authed := userGroup.Group("/") // 需要登录保护
		authed.Use(middleware.AuthMiddleware())
		{
			// 用户操作
			authed.PUT("update", user.UserUpdate)
			// 发送邮件
			authed.POST("sending-email", user.SendEmail)
		}
	}

	return r
}