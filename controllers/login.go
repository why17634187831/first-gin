package authUser

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// nickname := c.Query("nickname")
	// tel := c.Query("tel")
	nickname := c.Query("nickname")
	tel := c.PostForm("tel")
	// ctx.String(http.StatusOK, "这是后台登录页面")
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"test":    nickname,
		"tel":     tel,
	})
}
