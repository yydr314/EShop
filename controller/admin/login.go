package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(ctx *gin.Context) {
	ctx.String(http.StatusOK, "do login")
}
