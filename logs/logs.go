package logs

import (
	"strings"

	"github.com/astaxie/beego/logs"
)

var (
	log *logs.BeeLogger
)

func init() {
	log = logs.NewLogger()
	log.SetLogFuncCallDepth(3)
	log.SetLogger(logs.AdapterConsole)
	log.SetLogger(logs.AdapterFile, `{"filename":"logs/info.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	log.SetLogger(logs.AdapterFile, `{"filename":"logs/error.log","separate":["error"]}`)
	log.EnableFuncCallDepth(true)
	log.Debug("this is a debug message")

}

func Info(v ...interface{}) {
	log.Info(generateFmtStr(len(v)), v...)
}
func Error(v ...interface{}) {
	log.Error(generateFmtStr(len(v)), v...)
}
func Debug(v ...interface{}) {
	log.Debug(generateFmtStr(len(v)), v...)
}
func Warning(v ...interface{}) {
	log.Warning(generateFmtStr(len(v)), v...)
}
func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
