package routers

import (
	"demo/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	// 创建文件 设置输出日志文件路径
	if os.Getenv("runmode") == "prod" {
		utils.LogFileRecord()
	}
	r.Use(gin.Logger())
	return r
}
