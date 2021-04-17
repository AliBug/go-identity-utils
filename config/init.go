package config

import (
	"log"

	"github.com/spf13/viper"
)

// 设置 要从环境变量中要读取的 变量前缀（自动转为大写)
func init() {
	viper.SetEnvPrefix("config")

	// 绑定 配置文件名 环境变量
	viper.BindEnv("name")
	// 绑定 配置文件路径 环境变量
	viper.BindEnv("path")
	// 绑定 配置文件类型 环境变量
	viper.BindEnv("type")

	// 设置 配置文件名 缺省值
	viper.SetDefault("name", "config")
	// 设置 配置文件路径 缺省值

	// ⚠️ 下一次 把 缺省路径 放到 绝对路径 /conf
	viper.SetDefault("path", "/conf")
	// 设置 配置文件
	viper.SetDefault("type", "yaml")
	configFileName := viper.GetString("name")
	configFilePath := viper.GetString("path")
	configFileType := viper.GetString("type")

	// 确定要读取的配置文件
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)
	viper.AddConfigPath(configFilePath)

	// 处理读取错误
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config file error: %s", err)
	}
}
