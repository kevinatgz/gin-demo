package main

import (
	_ "gin-demo/config"
	db "gin-demo/database"
	//_ "gin-demo/database"
	"gin-demo/routes"
	"os"
)

func main() {
	r := routes.InitRouter()
	port := os.Getenv("HTTP_PORT")
	dbOpen :=db.InitDb()
	defer dbOpen.Close()
	//r := gin.Default()
	//r.GET("/ping2", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run(":" + port) // 监听并在 0.0.0.0:8080 上启动服务

}