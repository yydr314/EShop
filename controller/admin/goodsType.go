package admin

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type GoodsTypeController struct {
	BaseController
}

func (con GoodsTypeController) Index(ctx *gin.Context) {
	//	建立角色結構列表
	goodsTypeList := []models.GoodsType{}
	//	解析資料庫中的資料進結構
	models.DB.Find(&goodsTypeList)

	//	資料回傳至模板
	ctx.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsTypeController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}

func (con GoodsTypeController) DoAdd(ctx *gin.Context) {
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入的參數不正確", "/admin/goodsType/add")
		return
	}

	if title == "" {
		con.Error(ctx, "角色標題不能為空", "/admin/goodsType/add")
		return
	}

	goodsType := models.GoodsType{}

	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(time.Now().Unix())

	err = models.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(ctx, "增加商品種類失敗 請重試", "/admin/goodsType/add")
	} else {
		con.Success(ctx, "增加商品種類成功", "/admin/goodsType")
	}

}

func (con GoodsTypeController) Edit(ctx *gin.Context) {
	//	獲取Id轉換為數字類型
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/amdin/goodsType")
	}
	//	定義結構體要查的Id
	goodsType := models.GoodsType{Id: id}

	//	查詢資料
	models.DB.Find(&goodsType)

	ctx.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
		"goodsType": goodsType,
	})
}

func (con GoodsTypeController) DoEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		con.Error(ctx, "傳輸數據錯誤", "/admin/goodsType")
	}
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入參數錯誤", "/admin/goodsType")
	}

	if title == "" {
		con.Error(ctx, "角色標題不能為空", "/admin/goodsType/edit?id="+strconv.Itoa(id))
		return
	}

	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err = models.DB.Save(&goodsType).Error
	if err != nil {
		con.Error(ctx, "修改失敗", "/admin/goodsType/edit?id="+strconv.Itoa(id))
	} else {
		con.Success(ctx, "修改數據成功", "/admin/goodsType")
	}
}

func (con GoodsTypeController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Delete(&goodsType)
		con.Success(ctx, "刪除數據成功", "/admin/goodsType")
	}
}
