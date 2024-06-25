package router

import (
	"xiaomiShop/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters:=r.Group("/admin")
	{
		adminRouters.GET("/login",admin.LoginController{}.Index)
		adminRouters.POST("/doLogin",admin.LoginController{}.DoLogin)
	}

}