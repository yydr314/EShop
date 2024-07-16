package admin

import (
	"net/http"
	"strconv"
	"strings"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(ctx *gin.Context) {
	//	資料回傳至模板
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	ctx.HTML(http.StatusOK, "admin/access/index.html", gin.H{"accessList": accessList})
}

func (con AccessController) Add(ctx *gin.Context) {
	//	獲取頂級模塊
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	ctx.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(ctx *gin.Context) {
	moduleName := strings.Trim(ctx.PostForm("module_name"), " ")
	actionName := ctx.PostForm("action_name")
	accessType, err := strconv.Atoi(ctx.PostForm("type"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/add")
		return
	}
	url := ctx.PostForm("url")
	moduleId, err := strconv.Atoi(ctx.PostForm("module_id"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/add")
		return
	}
	sort, err := strconv.Atoi(ctx.PostForm("sort"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/add")
		return
	}
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/add")
		return
	}
	description := ctx.PostForm("description")

	if moduleName == "" {
		con.Error(ctx, "模塊名稱不能為空", "/admin/access/add")
	}

	access := models.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err = models.DB.Create(&access).Error
	if err != nil {
		con.Error(ctx, "資料添加失敗", "/admin/access/add")
		return
	}
	con.Success(ctx, "添加資料成功", "/admin/access/")
}
