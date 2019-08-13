package log

import (
	"github.com/astaxie/beego/logs"
)

var (
	log *logs.BeeLogger
)

func init() {
	log = logs.NewLogger()
	log.SetLogFuncCallDepth(3)
	log.SetLogger(logs.AdapterConsole)
	log.EnableFuncCallDepth(true)
	log.Debug("this is a debug message")
}

func Info(format string, v ...interface{}) {
	log.Info(format, v...)
}
func Error(format string, v ...interface{}) {
	log.Error(format, v...)
}
