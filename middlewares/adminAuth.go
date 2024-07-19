package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"xiaomiShop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
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
		//	json字串反解析回物件
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				ctx.Redirect(302, "/admin/login")
			}
		} else { //	用戶登入成功 權限判斷
			//	因為pathname是/admin/xxx/xxx，但在資料庫中儲存的是xxx/xxx，所以要將/admin/替換為空
			urlPath := strings.Replace(pathname, "/admin/", "", 1)

			//	用戶不是超級管理員，訪問的地址也不是排除地址
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				//	獲取當前角色的權限列表
				roleAccess := []models.RoleAccess{}
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]struct{})
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = struct{}{}
				}

				access := models.Access{}
				models.DB.Where("url=?", urlPath).Find(&access)

				if _, ok := roleAccessMap[access.Id]; !ok {
					ctx.String(http.StatusForbidden, "沒有權限")
					ctx.Abort()
				}
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

// 排除權限判斷的方法
func excludeAuthPath(urlPath string) bool {
	//從ini文件抓取排除地址
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	path := config.Section("").Key("excludeAuthPath").String()

	pathSlice := strings.Split(path, ",")

	for _, v := range pathSlice {
		if v == urlPath {
			return true
		}
	}
	return false

}
