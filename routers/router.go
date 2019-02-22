package routers

import (
	"github.com/TopJohn/go-web-framework/controller"
	"github.com/TopJohn/go-web-framework/middleware/jwt"
	"github.com/TopJohn/go-web-framework/middleware/logger"
	"github.com/TopJohn/go-web-framework/middleware/recovery"
	"github.com/TopJohn/go-web-framework/pkg/conf"
	"github.com/TopJohn/go-web-framework/pkg/pprof"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//wrap pprof
	pprof.Wrap(r)
	r.Use(logger.LoggerWithBeeLogger())
	r.Use(recovery.RecoveryWithBeeLogger())
	gin.SetMode(conf.ServerConfig.RunMode)
	r.StaticFS("/static", http.Dir(conf.AppConfig.StaticPath))

	//unauthorized api
	r.POST("/login", controller.Login)

	//authorized api
	v1 := r.Group("/v1").Use(jwt.JWT())
	{

	}
	return r
}
