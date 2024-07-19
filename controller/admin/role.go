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

func (con RoleController) Auth(ctx *gin.Context) {
	roleId, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/role")
		return
	}

	//	獲取權限列表
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	//	找到當前角色的所有權限
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Find(&roleAccess)

	//	將該角色的所有權限寫入map中
	roleAccessMap := make(map[int]struct{})
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = struct{}{}
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

	ctx.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})
}

func (con RoleController) DoAuth(ctx *gin.Context) {
	roleId, err := strconv.Atoi(ctx.PostForm("role_id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/role")
		return
	}
	//	獲取權限id切片
	accessIds := ctx.PostFormArray("access_node[]")

	//刪除當前角色對應的權限
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Delete(&roleAccess)

	//	增加當前角色對應的權限

	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := strconv.Atoi(v)
		roleAccess.AccessId = accessId
		err := models.DB.Create(&roleAccess).Error
		if err != nil {
			con.Error(ctx, "添加權限失敗", "/admin/role")
			return
		}
	}
	con.Success(ctx, "添加權限成功", "/admin/role")
}
