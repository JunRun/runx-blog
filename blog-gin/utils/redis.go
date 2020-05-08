/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/7 8:19 下午
 */
package utils

import (
	"github.com/go-redis/redis"

	"github.com/JunRun/blog-gin/conf"
)

const (
	CACHE = 1
)

type RedisPool struct {
}

func (r *RedisPool) Get() {

}
func NewRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Config.GetString("redis.host"),
		Password: "",
		DB:       conf.Config.GetInt("DB"),
	})
	if _, err := client.Ping().Result(); err != nil {
		return err
	}
	return nil
	//RedisConn := &redis.Pool{
	//}
}
