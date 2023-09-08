package utils

import (
	"github.com/gin-gonic/gin"
	"golang-restful-api-framework/internal/model/vo"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, vo.Response{
		0,
		data,
		"ok",
	})
}

func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, vo.Response{
		errorCode,
		nil,
		msg,
	})
}
