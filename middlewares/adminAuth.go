package middlewares

import (
	"encoding/json"
	"strings"
	"xiaomiShop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(ctx *gin.Context) {
	//	先獲取訪問的url
	pathname := strings.Split(ctx.Request.URL.String(), "?")[0] //	用split將captcha後面的隨機數去掉

	//	獲取session保存的訊息
	session := sessions.Default(ctx)
	userinfo := session.Get("userinfo")
	userinfoStr, ok := userinfo.(string)
	if ok {
		//	這裡判斷userinfo裡的訊息是否存在
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				ctx.Redirect(302, "/admin/login")
			}
		}
	} else {
		//	用戶沒有登入，代表沒有string類型
		//	要排除不需要權限驗證的路由
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			ctx.Redirect(302, "/admin/login")
		}
	}

}
