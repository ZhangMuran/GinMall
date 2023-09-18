package user

import (
	"net/http"

	"github.com/ZhangMuran/GinMall/service/user"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister user.UserRegisterService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}