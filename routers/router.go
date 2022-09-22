package routers

import (
	authUser "first-gin/controllers"

	"github.com/gin-gonic/gin"
)

func IndexInit(r *gin.Engine) {
	adminR := r.Group("/admin")
	adminR.POST("/login", authUser.Login)
}
