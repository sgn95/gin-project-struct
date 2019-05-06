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
	logFile, err := os.Create("logs/gin.log")
	if err != nil {

	}
	gin.DefaultWriter = io.MultiWriter(logFile)
	// 每天定时打包日志文件
	fileName := fmt.Sprintf("logs/%s-%s.log", setting.AppSetting["AppName"], time.Now().Format("2006-01-02"))
	// 每天打包日志文件
	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for range ticker.C {
			err = os.Rename("logs/gin.log", fileName)
		}
	}()
	if err != nil {
		fmt.Println(err)
	}
}
