package user

import (
	"net/http"

	"github.com/ZhangMuran/GinMall/service/user"
	"github.com/ZhangMuran/GinMall/service/utils"
	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	var userUpdate user.UserUpdateService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}