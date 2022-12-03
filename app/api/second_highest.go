package api

import (
	"api_test/app/service/no1796"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 1796. 字符串中第二大的数字
// @Description 给你一个混合字符串 s ，请你返回 s 中 第二大 的数字，如果不存在第二大的数字，请你返回 -1 。混合字符串 由小写英文字母和数字组成。
// @Produce  json
// @Param s formData string true "混合字符串 s"
// @Tags version
// @Produce json
// @Success 200 {string} object 字符串中第二大的数字或者-1
// @Router /api/no1796 [post]
func ApiSecondHighest(c *gin.Context) {
	s := c.PostForm("s")
	result := no1796.SecondHighest(s)
	c.JSON(http.StatusOK, result)
}
