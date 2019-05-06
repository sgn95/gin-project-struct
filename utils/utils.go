package utils

import (
	"demo/setting"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LogFileRecord 日志记录操作
func LogFileRecord() {
	// 先创建
	logFile, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile)
	// 每天定时打包日志文件
	fileName := fmt.Sprintf("%s-%s-.log", setting.AppSetting["AppName"], time.Now().Format("2006-01-02"))
	err := os.Rename("logs/gin.log", fileName)
	if err != nil {

	}
}
