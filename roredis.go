package roredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var rclient *redis.Client

const defaultPort = "6379"
const defaultHost = "localhost"

type RedisCfg struct {
	Host, Port, Password string
	DB                   int
}

// The internal redis client must be initialized before other functions called
func InitRedis(cfg RedisCfg) {
	var host, port string

	if cfg.Host == "" {
		host = defaultHost
	}
	if cfg.Port == "" {
		port = defaultPort
	}

	rclient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: cfg.Password,
		DB:       cfg.DB, // 0 happens to be the default DB
	})
}

func Ping() string {
	if rclient == nil {
		return ""
	}
	pong, err := rclient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis ping failed")
		return ""
	}
	return pong
	//	cl := Redis().LRange("testkey", 0, -1)
}

// Set expiration time to zero for no expiration
func Set(key, value string, expiration time.Duration) error {
	if rclient == nil {
		return errors.New("redis client not initialized - call InitRedis first")
	}
	return rclient.Set(context.Background(), key, value, expiration).Err()
}

func Get(key string) (val string, err error) {
	if rclient == nil {
		return val, errors.New("redis client not initialized - call InitRedis first")
	}

	val, err = rclient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", errors.New("Key does not exist")
	} else if err != nil {
		return "", err
	}
	return val, err
}

func Del(key string) error {
	if rclient == nil {
		return errors.New("redis client not initialized - call InitRedis first")
	}
	return rclient.Del(context.Background(), key).Err()
}
