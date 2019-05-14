package setting

import (
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

// AppSetting 全局配置项
var AppSetting map[string]string

// SetUp 配置文件 配置 默认读取内容都是string
func SetUp() {
	conf, _ := ini.Load("conf/app.ini")
	defValue, _ := conf.GetSection(ini.DEFAULT_SECTION)
	AppSetting = defValue.KeysHash()
	if os.Getenv("runmode") == "dev" {
		devSetting, _ := conf.GetSection("dev")
		devSettingHash := devSetting.KeysHash()
		for key, value := range devSettingHash {
			AppSetting[key] = value
		}
	}
	if os.Getenv("runmode") == "prod" {
		prodSetting, _ := conf.GetSection("dev")
		prodSettingHash := prodSetting.KeysHash()
		for key, value := range prodSettingHash {
			AppSetting[key] = value
		}
	}
}

// ToInt 读取配置文件 Int
func ToInt(confStr string) (int, error) {
	confInt, err := strconv.Atoi(AppSetting[confStr])
	return confInt, err
}

// ToString 读取配置文件 转成string类型
func ToString(confStr string) string {
	return AppSetting[confStr]
}
