package user

import (
	"net/http"

	"github.com/ZhangMuran/GinMall/service/user"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var UserLogin user.UserLoginService
	if err := c.ShouldBind(&UserLogin); err == nil {
		res := UserLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}