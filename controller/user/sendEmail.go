package user

import (
	"net/http"

	"github.com/ZhangMuran/GinMall/service/user"
	"github.com/ZhangMuran/GinMall/service/utils"
	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	var userUpdate user.UserEmailService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.SendEmail(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}