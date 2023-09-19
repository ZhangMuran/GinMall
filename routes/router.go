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
	}

	return r
}