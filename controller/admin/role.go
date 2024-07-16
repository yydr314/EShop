package admin

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(ctx *gin.Context) {
	//	建立角色結構列表
	roleList := []models.Role{}
	//	解析資料庫中的資料進結構
	models.DB.Find(&roleList)

	//	資料回傳至模板
	ctx.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(ctx *gin.Context) {
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")

	if title == "" {
		con.Error(ctx, "角色標題不能為空", "/admin/role/add")
		return
	}

	role := models.Role{}

	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(time.Now().Unix())

	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(ctx, "增加角色失敗 請重試", "/admin/role/add")
	} else {
		con.Success(ctx, "增加角色成功", "/admin/role")
	}

	ctx.String(http.StatusOK, "DoAdd")
}

func (con RoleController) Edit(ctx *gin.Context) {
	//	獲取Id轉換為數字類型
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/amdin/role")
	}
	//	定義結構體要查的Id
	role := models.Role{Id: id}

	//	查詢資料
	models.DB.Find(&role)

	ctx.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})
}

func (con RoleController) DoEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		con.Error(ctx, "傳輸數據錯誤", "/admin/role")
	}
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")

	if title == "" {
		con.Error(ctx, "角色標題不能為空", "/admin/role/edit")
		return
	}

	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description

	err = models.DB.Save(&role).Error
	if err != nil {
		con.Error(ctx, "修改失敗", "/admin/role/edit?id="+strconv.Itoa(id))
	} else {
		con.Success(ctx, "修改數據成功", "/admin/role/role")
	}
}

func (con RoleController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(ctx, "刪除數據成功", "/admin/role")
	}
}
