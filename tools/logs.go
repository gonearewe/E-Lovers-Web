package tools

import (
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

func NewLogger() *logs.BeeLogger {

	//example showing how to use beego/logs
	//init process
	//later you may use mylog.Debug("information") and so on to record
	//mylog.Flush() may be calld to flush buf into output manually
	//mylog.Close() should be called at last to compete the whole process
	mylog := logs.NewLogger(10000)

	jsonConfig := `{
        "filename" : "test.log", 
        "maxlines" : 1000,       
        "maxsize"  : 10240       
	}`
	date := time.Now().Format("2006-01-02")
	filename := date + ".log"
	strings.ReplaceAll(jsonConfig, "test.log", filename)

	mylog.SetLogger("file", jsonConfig)
	mylog.SetLevel(logs.LevelDebug) // set the level above which will be recorded
	mylog.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

	return mylog

}
