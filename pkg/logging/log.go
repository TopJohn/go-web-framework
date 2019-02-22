package logging

import "github.com/astaxie/beego/logs"

func Setup() {
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterMultiFile,
		`{
					"filename":"logs/log.log",
					"maxLines":1000000,
					"maxsize":1<<28,
					"daily":true,
					"maxDays":7,
					"rotate":true,
				  	"perm":0600,
					"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
				}`)
	logs.Async(1e3)
}
