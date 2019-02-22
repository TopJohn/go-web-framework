package main

import (
	"fmt"
	"github.com/TopJohn/go-web-framework/models"
	"github.com/TopJohn/go-web-framework/pkg/conf"
	"github.com/TopJohn/go-web-framework/pkg/logging"
	"github.com/TopJohn/go-web-framework/routers"
	"github.com/astaxie/beego/logs"
	"log"
	"net/http"
)

func main() {
	logging.Setup()
	models.Setup()
	router := routers.InitRouter()
	readTimeout := conf.ServerConfig.ReadTimeout
	writeTimeout := conf.ServerConfig.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)
	//Default maxHeaderBytes is 1MB
	maxHeaderBytes := http.DefaultMaxHeaderBytes

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()

	if err != nil {
		logs.Error("server.ListenAndServe() error: %v", err)
	}

}
