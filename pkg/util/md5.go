package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/TopJohn/go-web-framework/pkg/conf"
)

func MD5(value string) string {
	md5 := md5.New()
	md5.Write([]byte(value))
	md5.Write([]byte(conf.AppConfig.Salt))
	result := md5.Sum(nil)
	return hex.EncodeToString(result)
}
