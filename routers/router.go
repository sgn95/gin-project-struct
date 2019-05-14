package routers

import (
	"demo/pkg/logs"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(logs.Logger())
	r.Use(gin.Recovery())
	return r
}
