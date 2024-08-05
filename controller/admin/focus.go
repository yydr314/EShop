package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(ctx *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	ctx.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(ctx *gin.Context) {
	title := ctx.PostForm("title")
	focusType, err := strconv.Atoi(ctx.PostForm("focus_type"))
	if err != nil {
		con.Error(ctx, "輸入資料錯誤", "/admin/focus/add")
	}

	link := ctx.PostForm("link")
	sort, err := strconv.Atoi(ctx.PostForm("sort"))
	if err != nil {
		con.Error(ctx, "請輸入正確的排序值", "/admin/focus/add")
	}
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "輸入資料錯誤", "/admin/focus/add")
	}

	focusImgSrc, err := models.UploadFile(ctx, "focus_img")
	if err != nil {
		fmt.Println(err)
	}

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(time.Now().Unix()),
	}
	err = models.DB.Create(&focus).Error
	if err != nil {
		con.Error(ctx, "上傳失敗", "/admin/focus/add")
	} else {
		con.Success(ctx, "新增成功！", "/admin/focus")
	}

}

func (con FocusController) Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入參數錯誤", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	ctx.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		con.Error(ctx, "非法傳輸", "/admin/focus")
	}
	title := ctx.PostForm("title")
	focusType, err := strconv.Atoi(ctx.PostForm("focus_type"))
	if err != nil {
		con.Error(ctx, "輸入資料錯誤", "/admin/focus/edit?id="+strconv.Itoa(id))
	}

	link := ctx.PostForm("link")
	sort, err := strconv.Atoi(ctx.PostForm("sort"))
	if err != nil {
		con.Error(ctx, "請輸入正確的排序值", "/admin/focus/edit?id="+strconv.Itoa(id))
	}
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "輸入資料錯誤", "/admin/focus/edit?id="+strconv.Itoa(id))
	}

	focusImgSrc, err := models.UploadFile(ctx, "focus_img")
	if err != nil {
		fmt.Println(err)
	}

	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status

	if focusImgSrc != "" {
		focus.FocusImg = focusImgSrc
	}
	err = models.DB.Save(&focus).Error
	if err != nil {
		con.Error(ctx, "修改數據失敗", "/admin/focus/edit?id="+strconv.Itoa(id))
	} else {
		con.Success(ctx, "修改數據成功", "/admin/focus")
	}
}

func (con FocusController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/focus")
	} else {
		focus := models.Focus{Id: id}
		models.DB.Delete(&focus)
		
		//	這裡根據需求看要不要刪除圖片，很多系統只是在資料庫中標記刪除並不會把檔案真的刪掉
		//	os.Remove(focus.FocusImg)
		con.Success(ctx, "刪除數據成功", "/admin/focus")
	}
}
