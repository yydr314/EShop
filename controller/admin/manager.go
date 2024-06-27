package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})
}

func (con ManagerController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{})
}

func (con ManagerController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}

func(con ManagerController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete")
}