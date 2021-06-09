package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const redisURLTemplate = "redis://:%s@%s/%s"

/*
// ReadRedisConfig - return redis conn url
func ReadRedisConfig(sect string) string {
	// 读取 redis 参数
	redisHost := viper.GetString(fmt.Sprintf("%s.host", sect))
	redisPass := viper.GetString(fmt.Sprintf("%s.pass", sect))
	redisDbName := viper.GetString(fmt.Sprintf("%s.database", sect))
	// ⚠️ pass 中 有 @ 会出错
	return fmt.Sprintf(redisURLTemplate, redisPass, redisHost, redisDbName)
}
*/

// RedisConfig - contain redis config
type RedisConfig interface {
	GetAddress() string
	GetPassword() string
	GetDB() int
}

type redisConfig struct {
	address  string
	password string
	db       int
}

func (r *redisConfig) GetAddress() string {
	return r.address
}
func (r *redisConfig) GetPassword() string {
	return r.password
}
func (r *redisConfig) GetDB() int {
	return r.db
}

// ReadRedisConfig - return RedisConfig interface
func ReadRedisConfig(sect string) RedisConfig {
	// 读取 redis 参数
	redisHost := viper.GetString(fmt.Sprintf("%s.host", sect))
	redisPort := viper.GetString(fmt.Sprintf("%s.port", sect))
	redisPass := viper.GetString(fmt.Sprintf("%s.pass", sect))
	redisDbName := viper.GetInt(fmt.Sprintf("%s.database", sect))

	config := &redisConfig{db: redisDbName, password: redisPass}

	if redisPort != "" {
		config.address = fmt.Sprintf("%s:%s", redisHost, redisPort)
	} else {
		config.address = fmt.Sprintf("%s:%s", redisHost, "6379")
	}
	return config
}
