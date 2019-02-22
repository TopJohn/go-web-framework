package util

import (
	"github.com/TopJohn/go-web-framework/pkg/conf"
	"github.com/gin-gonic/gin"
)
import "github.com/Unknwon/com"

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * conf.AppConfig.PageSize
	}
	return result
}
