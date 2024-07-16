package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xiaomiShop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(ctx *gin.Context) {
	verifyValue := ctx.PostForm("verifyValue")
	captchaId := ctx.PostForm("captchaId")

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//	先驗證驗證碼的正確性
	if ok := models.VerifyCaptcha(captchaId, verifyValue); ok {
		//	查詢資料庫判斷用戶以及密碼是否存在
		userinfoList := []models.Manager{}
		password = models.Md5(password)
		models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

		if len(userinfoList) > 0 {
			//	這裡執行登入 保存用戶信息後執行跳轉（這裡用session存）
			session := sessions.Default(ctx)
			//	這裡因為session無法直接保存結構體對應的切片，因此把結構轉為json字串
			userinfoSlice, _ := json.Marshal(userinfoList)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()

			con.Success(ctx, "登入成功", "/admin")
		} else {
			con.Error(ctx, "用戶名或密碼錯誤", "/admin/login")
		}
	} else {
		con.Error(ctx, "驗證碼驗證失敗", "/admin/login")
	}

}

func (con LoginController) Captcha(ctx *gin.Context) {
	id, b64s, ans, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId":     id,
		"captchaImage":  b64s,
		"captchaAnswer": ans,
	})
}

func (con LoginController) LoginOut(ctx *gin.Context) {
	//	先銷毀session
	session := sessions.Default(ctx)
	session.Delete("usderinfo")
	session.Save()
	con.Success(ctx, "退出登入成功", "/admin/login")
}
