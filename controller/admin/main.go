package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"xiaomiShop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		//	Preload傳入函數的原因是，我們在預加載的時候希望預加載的資料是排序好的，所以要這樣寫
		//	用DESC則sort值越大越上面
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order(("access.sort DESC"))
		}).Order("sort DESC").Find(&accessList)

		//	找到當前角色的所有權限
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)

		//	將該角色的所有權限寫入map中
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		//	判斷所有權限中的哪些權限，在這個角色的權限map裡面
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
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		ctx.Redirect(302, "/admin/login")
	}

}

func (con MainController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共修改狀態的方法
func (con MainController) ChangeStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "傳入參數錯誤",
		})
		return
	}

	table := ctx.Query("table")
	field := ctx.Query("field")

	/*
		利用ABS函數（絕對值）來將0變1,1變0，因為：
		status = ABS(0 - 1) = 1
		status = ABS(1 - 1) = 0
	*/
	err = models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "修改數據失敗",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改數據成功",
	})
}
