package models

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 鏈接資料庫
func init() {
	//	讀取.ini內的資料庫配置文件
	config, err := ini.Load("./conf/app.ini") //	這裡路徑寫這樣是因為到時候引入也是透過main.go，所以不需要返回根目錄
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//	指派.ini中的資料
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

	// dsn := "root:Willie81926@tcp(127.0.0.1:3306)/gin"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
