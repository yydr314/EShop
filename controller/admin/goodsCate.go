package admin

import (
	"net/http"
	"strconv"
	"time"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type GoodsCateController struct {
	BaseController
}

func (con GoodsCateController) Index(ctx *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)

}

func (con GoodsCateController) Add(ctx *gin.Context) {
	//	加載頂級分類
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&goodsCateList)
	ctx.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoAdd(ctx *gin.Context) {
	title := ctx.PostForm("title")
	pid, err := strconv.Atoi(ctx.PostForm("pid"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/add")
		return
	}
	link := ctx.PostForm("link")
	template := ctx.PostForm("template")
	subTitle := ctx.PostForm("sub_title")
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	sort, err := strconv.Atoi(ctx.PostForm("sort"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/add")
		return
	}
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/add")
		return
	}

	//	添加分類有可能不需要圖片，所以不接收 err
	cateImgDir, _ := models.UploadFile(ctx, "cate_img")

	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		Link:        link,
		Template:    template,
		SubTitle:    subTitle,
		Keywords:    keywords,
		Description: description,
		Sort:        sort,
		Status:      status,
		CateImg:     cateImgDir,
		AddTime:     int(time.Now().Unix()),
	}

	err = models.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(ctx, "增加商品失敗", "/admin/goodsCate/add")
		return
	} else {
		con.Success(ctx, "增加商品成功", "/admin/goodsCate")
	}
}

func (con GoodsCateController) Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入參數錯誤", "/admin/goodsCate")
	}

	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)

	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Find(&goodsCateList)

	ctx.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/edit")
	}
	title := ctx.PostForm("title")
	pid, err := strconv.Atoi(ctx.PostForm("pid"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	link := ctx.PostForm("link")
	template := ctx.PostForm("template")
	subTitle := ctx.PostForm("sub_title")
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	sort, err := strconv.Atoi(ctx.PostForm("sort"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入參數不正確", "/admin/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}

	//	添加分類有可能不需要圖片，所以不接收 err
	cateImgDir, _ := models.UploadFile(ctx, "cate_img")

	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)

	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status

	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}

	err = models.DB.Save(&goodsCate).Error
	if err != nil {
		con.Error(ctx, "修改資料失敗，請重試", "/admin/goodsCate/edit?id="+strconv.Itoa(id))
	} else {
		con.Success(ctx, "修修改資料成功", "/admin/goodsCate")
	}
}

func (con GoodsCateController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "傳入數據錯誤", "/admin/goodsCate")
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)

	if goodsCate.Pid == 0 {
		goodsCateList := []models.GoodsCate{}
		models.DB.Where("pid = ?", goodsCate.Pid).Find(&goodsCateList)

		if len(goodsCateList) > 0 {
			con.Error(ctx, "請先刪除商品再刪除頂級分類", "/admin/goodsCate")
			return
		} else {
			models.DB.Delete(&goodsCate)
			con.Success(ctx, "刪除成功", "/admin/goodsCate")
		}
	} else {
		models.DB.Delete(&goodsCate)
		con.Success(ctx, "刪除成功", "/admin/goodsCate")
	}
}
