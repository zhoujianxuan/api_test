package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AppVersion godoc
// @Summary 测试接口
// @Schemes
// @Description 这是一个返回version数据并且实时的服务器时间的接口
// @Tags version
// @Accept json
// @Produce json
// @Success 200 {string} object 获取版本
// @Router /api/version [get]
func AppVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v1.0.0",
		"time":    time.Now(),
	})
}
