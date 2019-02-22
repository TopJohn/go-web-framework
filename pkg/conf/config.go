package conf

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/ini.v1"
	"time"
)

type App struct {
	JwtSecret  string
	PageSize   int
	PrefixUrl  string
	StaticPath string
	Salt       string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	User        string
	Password    string
	Host        string
	Port        int
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var AppConfig = &App{}
var ServerConfig = &Server{}
var DatabaseConfig = &Database{}
var RedisConfig = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		logs.Error("Config.Setup, fail to parse 'conf/conf.ini':%v", err)
	}
	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
	mapTo("database", DatabaseConfig)
	mapTo("redis", RedisConfig)

	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	RedisConfig.IdleTimeout = RedisConfig.IdleTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logs.Error("Cfg.MapTo RedisConfig err: %v", err)
	}
}
