package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const redisURLTemplate = "redis://:%s@%s/%s"

// ReadRedisConfig - return redis conn url
func ReadRedisConfig(sect string) string {
	// 读取 redis 参数
	redisHost := viper.GetString(fmt.Sprintf("%s.host", sect))
	redisPass := viper.GetString(fmt.Sprintf("%s.pass", sect))
	redisDbName := viper.GetString(fmt.Sprintf("%s.database", sect))
	// ⚠️ pass 中 有 @ 会出错
	return fmt.Sprintf(redisURLTemplate, redisPass, redisHost, redisDbName)
}
