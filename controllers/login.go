package authUser

import (
	"github.com/gin-gonic/gin"
)

type loginParam struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var person loginParam
	if err := c.ShouldBind(&person); err != nil {
		// c.String(500, fmt.Sprint(err))
		c.JSON(200, gin.H{
			"success": true,
			"code":    500,
			"message": err,
		})
		return
	}
	// c.JSON(200, fmt.Sprintf("%#v", person))
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"test":    person,
	})
	//登录
	// nickname := c.Query("nickname")
	// tel := c.PostForm("tel")
	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"code":    200,
	// 	"test":    nickname,
	// 	"tel":     tel,
	// })
}
