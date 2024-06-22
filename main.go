package main

import (
	"xiaomiShop/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitSQLRouter(r)

	r.Run()
}
