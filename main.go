package main

import (
	"fmt"
	"os"
	"xiaomiShop/router"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	r := gin.Default()

	router.InitSQLRouter(r)
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
