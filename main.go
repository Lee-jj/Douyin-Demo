package main

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"DOUYIN-DEMO/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// go service.RunMessageServer()

	r := gin.Default()

	dao.InitMySQL()
	dao.DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Favorite{})

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
