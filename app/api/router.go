package api

import "github.com/gin-gonic/gin"

func Init(e *gin.Engine) {
	router := e.Group("/api")
	router.GET("/version", AppVersion)
	router.POST("/no1796", ApiSecondHighest)
}
