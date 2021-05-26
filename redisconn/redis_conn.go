package redisconn

import (
	"fmt"

	"github.com/alibug/go-identity-utils/config"
	"github.com/go-redis/redis/v8"
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
	conf := config.ReadRedisConfig(sect)
	opt := &redis.Options{Addr: conf.GetAddress(), Password: conf.GetPassword(), DB: conf.GetDB()}
	return redis.NewClient(opt)
}
