/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/7 8:19 下午
 */
package utils

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/JunRun/blog-gin/conf"
)

const (
	CACHE = 1
)

var RedisClient *redis.Client

func NewRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        conf.Config.GetString("redis.host"),
		Password:    conf.Config.GetString("redis.password"),
		DB:          conf.Config.GetInt("redis.DB"),
		IdleTimeout: conf.Config.GetDuration("redis.IdleTimeout"),
	})
	if _, err := RedisClient.Ping().Result(); err != nil {
		fmt.Println("连接redis 错误 ", err)
	}
	//RedisConn := &redis.Pool{
	//}
}
