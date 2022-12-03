package main

import (
	"api_test/app/api"
	"api_test/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 结合算法题的一套api接口
// @version 1.0
// @description Go 语言编写
func main() {
	router := gin.Default()

	api.Init(router)

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8000")
}
