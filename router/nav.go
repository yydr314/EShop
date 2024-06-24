package router

import (
	"xiaomiShop/controller"

	"github.com/gin-gonic/gin"
)

func InitNavRouter(r *gin.Engine) {
	controller:=controller.NavController{}

	r.GET("/nav", controller.Index)
}