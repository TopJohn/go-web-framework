package conf

import (
	"github.com/go-ini/ini"
	"time"
)

var (
	Cfg     *ini.File
	RunMode string

	HttpPort int

	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	PageSize int

	JwtSecret string
)

func init() {
	var err error
	Cfg, err := ini.Load("conf/conf.ini")

}
