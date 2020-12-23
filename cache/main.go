package cache

import (
	"auth-gateway/util/log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Panic("连接Redis不成功", err)
	}

	RedisClient = client
}

// AppProxy 代理缓存
var AppProxy *gin.RouterGroup
