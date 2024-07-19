package admin

import (
	"encoding/json"
	"net/http"
	"xiaomiShop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(ctx *gin.Context) {
	//在header顯示用戶名
	//獲取userinfo 對應的session
	session := sessions.Default(ctx)
	userinfo := session.Get("userinfo")
	//類型斷言判斷userinfo是否為string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//	獲取用戶訊息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		//	獲取所有權限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

		//	找到當前角色的所有權限
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)

		//	將該角色的所有權限寫入map中
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		//	判斷所有權限有沒有在這個角色的權限map裡面
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		ctx.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"is_super":   userinfoStruct[0].IsSuper,
		})
	} else {
		ctx.Redirect(302, "/admin/login")
	}

}

func (con MainController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
