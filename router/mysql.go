package router

import (
	"xiaomiShop/controller"

	"github.com/gin-gonic/gin"
)

func InitSQLRouter(r *gin.Engine) {
	controller := controller.SQLController{}

	r.GET("/get", controller.GetInfo)
	r.GET("/new", controller.NewInfo)
	r.GET("/delete", controller.DeleteInfo)
	r.GET("/update", controller.UpdateInfo)
}
