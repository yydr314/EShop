package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (con BaseController) Success(ctx *gin.Context, message string, redirectUrl string) {
	ctx.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
func (con BaseController) Error(ctx *gin.Context, message string, redirectUrl string) {
	ctx.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
