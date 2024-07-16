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

func (con AccessController) Edit(ctx *gin.Context) {
	//	獲取要修改的數據
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "參數錯誤", "/admin/access")
	}
	access := models.Access{Id: id}

	models.DB.Find(&access)

	//	獲取頂級模塊
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	ctx.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
		"access":     access,
	})
}

func (con AccessController) DoEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		con.Error(ctx, "參數錯誤", "/admin/access/edit")
	}
	moduleName := strings.Trim(ctx.PostForm("module_name"), " ")
	actionName := ctx.PostForm("action_name")
	accessType, err := strconv.Atoi(ctx.PostForm("type"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/edit")
		return
	}
	url := ctx.PostForm("url")
	moduleId, err := strconv.Atoi(ctx.PostForm("module_id"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/edit")
		return
	}
	sort, err := strconv.Atoi(ctx.PostForm("sort"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/edit")
		return
	}
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入的參數錯誤", "/admin/access/edit")
		return
	}
	description := ctx.PostForm("description")

	if moduleName == "" {
		con.Error(ctx, "模塊名稱不能為空", "/admin/access/edit?id="+strconv.Itoa(id))
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status

	err = models.DB.Save(&access).Error
	if err != nil {
		con.Error(ctx, "修改數據失敗", "/admin/access")
		return
	}
	con.Success(ctx, "修改數據成功", "/admin/access")

}

func (con AccessController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/access")
	} else {
		//	獲取要刪除的數據
		access := models.Access{Id: id}
		models.DB.Find(&access)
		if access.ModuleId == 0 { //	表示access是頂級模塊
			accessList := []models.Access{}
			//	找到其他模塊是否屬於當前這個頂級模塊
			models.DB.Where("module_id=?", access.Id).Find(&accessList)

			if len(accessList) > 0 {
				con.Error(ctx, "當前模塊下方有菜單或操作，請刪除後再刪除此數據", "/admin/access")
				return
			}
			models.DB.Delete(&access)
			con.Success(ctx, "刪除數據成功", "/admin/access")
		} else { //	表示access是操作或菜單
			models.DB.Delete(&access)
			con.Success(ctx, "刪除數據成功", "/admin/access")
		}
	}
}
