package models

import (
	"demo/setting"

	"github.com/jinzhu/gorm"
	// 加载mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// SetUp 数据库初始化
func SetUp() {
	var err error
	db, err := gorm.Open(setting.AppSetting["DatabaseType"], setting.AppSetting["SQLConn"])
	if err != nil {

	}
	// 全局禁用表名复数
	db.SingularTable(true)
	// 设置最大连接数
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(8)
}
