package main

import (
	"fmt"
	"os"
	"text/template"
	"xiaomiShop/models"
	"xiaomiShop/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	r := gin.Default()

	//	自定義的模板函數，必須要放在加載模板前面
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	r.LoadHTMLGlob("templates/**/**/*")
	//	靜態網頁目錄
	r.Static("/static", "./static")

	//	創建cookie的儲存引擎，secret111是用來加密的秘鑰
	store := cookie.NewStore([]byte("secret111"))
	//	配置session中間件 store是前面創建好的儲存引擎，也可以換成其他的
	r.Use(sessions.Sessions("userinfo", store))
	/*
		這裡要注意中間件的順序，必須要先使用中間件，再來初始化路由
		否則會造成中間件失效
	*/

	router.AdminRoutersInit(r)
	router.InitNavRouter(r)

	//	使用ini模塊
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	//	獲取ini資料
	fmt.Println(cfg.Section("").Key("app_name").String())
	//	Seciton為[]內的文字，key為變數名。如果不屬於任意Section則Section傳入空
	fmt.Println(cfg.Section("mysql").Key("password").String())

	//	寫入ini數據
	// cfg.Section("").Key("app_name").SetValue("change test")
	// cfg.SaveTo("./conf/app.ini")
	r.Run()
}
