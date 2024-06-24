package controller

import (
	"fmt"
	"net/http"
	"time"
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
	user := models.User{
		Username: "itying",
		Age:      22,
		Email:    "kollug1548@gmail.com",
		AddTime:  int(time.Now().Unix()),
	}

	//	新增數據
	models.DB.Create(&user)

	fmt.Println(user)

	ctx.String(http.StatusOK, "新增成功")
}

func (SQLController) UpdateInfo(ctx *gin.Context) {
	//	設定查找條件
	user := models.User{Id: 6}
	models.DB.Find(&user)
	name := models.User{}.TableName()
	fmt.Println(name)
	//	也可以這樣設定結合上面兩個
	//	models.DB.Where("id = ?", 6).Find(&user)

	//	修改找到的文件內容
	user.Username = "我是修改過後的內容"
	user.Email = "testmail@gmail.com"
	//	更新資料用Save
	models.DB.Save(&user)

	ctx.String(http.StatusOK, "修改成功")
}

func (SQLController) DeleteInfo(ctx *gin.Context) {
	user := models.User{Id: 6}
	models.DB.Delete(&user)

	ctx.String(http.StatusOK, "刪除成功")
}
