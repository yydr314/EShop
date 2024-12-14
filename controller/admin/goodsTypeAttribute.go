package admin

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"xiaomiShop/models"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(ctx *gin.Context) {
	cateId:=ctx.Query("cate_id")
	//	資料回傳至模板
	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{})
}

func (con GoodsTypeAttributeController) DoAdd(ctx *gin.Context) {
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		con.Error(ctx, "傳入的參數不正確", "/admin/goodsTypeAttribute/add")
		return
	}

	if title == "" {
		con.Error(ctx, "角色標題不能為空", "/admin/goodsTypeAttribute/add")
		return
	}

	goodsType := models.GoodsType{}

	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(time.Now().Unix())

	err = models.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(ctx, "增加商品種類失敗 請重試", "/admin/goodsTypeAttribute/add")
	} else {
		con.Success(ctx, "增加商品種類成功", "/admin/goodsTypeAttribute")
	}

}
