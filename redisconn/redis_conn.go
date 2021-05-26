package redisconn

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// NewConn - 初始化 redis 连接
func NewConn(url string) (*redis.Client, error) {
	redisOption, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("Redis url parse error")
	}
	redisClient := redis.NewClient(redisOption)
	return redisClient, nil
}

// NewConnFromConfig 从配置文件中读取信息，返回 redis 客户端
func NewConnFromConfig(sect string) *redis.Client {
	redisHost := viper.GetString(fmt.Sprintf("%s.host", sect))
	redisPort := viper.GetString(fmt.Sprintf("%s.port", sect))
	redisPass := viper.GetString(fmt.Sprintf("%s.pass", sect))
	redisDbName := viper.GetInt(fmt.Sprintf("%s.database", sect))
	if redisPort == "" {
		redisPort = "6379"
	}
	redisAddress := fmt.Sprintf("%s:%s", redisHost, redisPort)
	opt := &redis.Options{Addr: redisAddress, Password: redisPass, DB: redisDbName}
	return redis.NewClient(opt)
}
