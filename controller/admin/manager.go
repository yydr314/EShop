package admin

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(ctx *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(ctx *gin.Context) {
	//	獲取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(ctx *gin.Context) {
	roleId, err := strconv.Atoi(ctx.PostForm("role_id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/manager")
		return
	}

	//	抓取前端傳入的資料
	username := strings.Trim(ctx.PostForm("username"), " ")
	password := strings.Trim(ctx.PostForm("password"), " ")
	email := strings.Trim(ctx.PostForm("email"), " ")
	mobile := strings.Trim(ctx.PostForm("mobile"), " ")

	//	判斷資料合法性
	if len(username) < 2 || len(password) < 6 {
		con.Error(ctx, "用戶名或密碼長度不合法", "/admin/manager/add")
		return
	}

	//判斷管理員是否存在
	managerLlist := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerLlist)
	if len(managerLlist) > 0 {
		con.Error(ctx, "管理員名稱已存在", "/admin/manager/add")
		return
	}

	//	執行增加管理員
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(time.Now().Unix()),
	}

	err = models.DB.Create(&manager).Error
	if err != nil {
		con.Error(ctx, "增加管理員失敗", "/admin/manager/add")
		return
	}

	con.Success(ctx, "增加管理員成功", "/admin/manager")
}

func (con ManagerController) Edit(ctx *gin.Context) {
	//	獲取管理員
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/manager")
		return
	}

	manager := models.Manager{Id: id}
	models.DB.Find(&manager)

	//	獲取所有角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/amdin/manager")
		return
	}
	roleId, err := strconv.Atoi(ctx.PostForm("role_id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/amdin/manager")
		return
	}

	username := strings.Trim(ctx.PostForm("username"), " ")
	password := strings.Trim(ctx.PostForm("password"), " ")
	email := strings.Trim(ctx.PostForm("email"), " ")
	mobile := strings.Trim(ctx.PostForm("mobile"), " ")

	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId

	if password != "" {
		//	判斷密碼長度是否合法
		if len(password) < 6 {
			con.Error(ctx, "密碼長度不合法", "/admin/manager/edit?id="+strconv.Itoa(id))
		}
		manager.Password = models.Md5(password)
	}
	err = models.DB.Save(&manager).Error
	if err != nil {
		con.Error(ctx, "修改數據失敗", "/admin/manager/edit?id="+strconv.Itoa(id))
	}
	con.Success(ctx, "修改數據成功", "/admin/manager")
}

func (con ManagerController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/manager")
	} else {
		manager := models.Manager{Id: id}
		models.DB.Delete(&manager)
		con.Success(ctx, "刪除數據成功", "/admin/manager")
	}
}
