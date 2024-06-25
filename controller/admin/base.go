package admin

import (
	"net/http"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (con BaseController) Index(ctx *gin.Context) {
	//	獲取nav數據
	navList := []models.Nav{}
	navInList := []models.Nav{}
	navSelectList := []models.Nav{}
	navOrderList := []models.Nav{}
	//	注意條件寫法
	models.DB.Where("id>3 AND id<9").Find(&navList)
	models.DB.Where("id IN ?", []int{3, 5, 6}).Find(&navInList) //	In用切片表示條件範圍
	models.DB.Select("id,title").Find(&navSelectList)           //	查到的數據只返回特定column
	models.DB.Order("sort asc").Find(&navOrderList)             //	用Order來排序找到的結果
	//	上面的東西可以在.Find後面加上.Count來計算數據的數量，要傳入int64類型

	//	可以用Exec來執行刪除修改新增SQL語法
	//models.DB.Exec("delete from user where id=?", 5)

	ctx.JSON(http.StatusOK, gin.H{
		"navList":       navList,
		"navInList":     navInList,
		"navSelectList": navSelectList,
		"navOrderList":  navOrderList,
	})
}
