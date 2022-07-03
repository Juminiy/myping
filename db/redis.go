package db

import (
	"context"
	"fmt"
	"time"
	config "github.com/Juminiy/myping/config"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

const (
	REDIS_DEFAULT_PING           = "v0:ping"
	REDIS_DEFAULT_PONG           = "v0:pong"
	REDIS_DEFAULT_TTL            = 1
	PING_RECORD_DEFAULT_PREFIX   = "v1:"
	PING_RECORD_DEFAULT_PREFIX_X = "v2:"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func init() {
	config.AppConfig()
	RedisConnect()
}

func RedisConnect() {
	//fmt.Println(viper.GetString("db.redis.addr"),viper.GetString("db.redis.password"),viper.GetInt("db.redis.db"))
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("db.redis.addr"),
		Password: viper.GetString("db.redis.password"),
		DB:       viper.GetInt("db.redis.db"),
	})
	// log error
	//fmt.Println(RedisClient.Options().Addr, RedisClient.Options().Password, RedisClient.Options().DB)
	// don't get the config
	// KeepTTL = -1
	err := RedisClient.Set(Ctx, REDIS_DEFAULT_PING, REDIS_DEFAULT_PONG, REDIS_DEFAULT_TTL*time.Second).Err()
	if err != nil {
		fmt.Println("😴 Your redis server may not be available, please check.")
		panic(err)
	}
}
