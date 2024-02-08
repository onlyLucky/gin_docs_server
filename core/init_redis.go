package core

import (
	"context"
	"gin_docs_server/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis(db ...int) *redis.Client {
	
	var redisDB int = 0
	if len(db)>0{
		redisDB = db[0]
	}
	redisConf := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr: redisConf.Addr(),
		Password: redisConf.Password,
		DB: redisDB,
		PoolSize: redisConf.PoolSize,
	})
	// 测试连接 ping
	_, cancel := context.WithTimeout(context.Background(),500*time.Millisecond)
	defer cancel()
	_, err:=client.Ping().Result()
	if err != nil {
		logrus.Fatalf("%s redis 连接失败 err: %s", redisConf.Addr(),err.Error())
	}
	return client
}