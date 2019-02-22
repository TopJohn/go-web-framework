package jwt

import (
	e "github.com/TopJohn/go-web-framework/pkg/error"
	"github.com/TopJohn/go-web-framework/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var err error
		code = e.SUCCESS
		// Bearer token
		token := c.GetHeader("Authorization")[7:]
		if token == "" {
			code = e.ERROR_AUTH_INVALID
		} else {
			_, err := jwt.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_INVALID
			} else {
			}

		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsgWithError(code, err),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
