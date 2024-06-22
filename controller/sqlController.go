package controller

import (
	"net/http"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type SQLController struct {
}

func (SQLController) GetInfo(ctx *gin.Context) {
	userList := []models.User{}

	//	查詢age>20的資料，用where後面寫字串條件
	models.DB.Where("age>20").Find(&userList)

	ctx.JSON(http.StatusOK, userList)
}

func (SQLController) NewInfo(ctx *gin.Context) {
	
}

func (SQLController) UpdateInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Put")
}

func (SQLController) DeleteInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Delete")
}
