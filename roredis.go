package roredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var RClient *redis.Client

const defaultPort = "6379"
const defaultHost = "localhost"

type RedisCfg struct {
	Host, Port, Password string
	DB                   int
}

// The internal redis client must be initialized before other functions called
func InitRedis(cfg RedisCfg) *redis.Client {
	var host, port string

	if cfg.Host == "" {
		host = defaultHost
	}
	if cfg.Port == "" {
		port = defaultPort
	}

	RClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: cfg.Password,
		DB:       cfg.DB, // 0 happens to be the default DB
	})

	return RClient
}

func Ping() string {
	if RClient == nil {
		return ""
	}
	pong, err := RClient.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println("Redis ping failed")
		return ""
	}
	return pong
	//	cl := Redis().LRange("testkey", 0, -1)
}

// Set expiration time to zero for no expiration
func Set(key, value string, expiration time.Duration) error {
	if RClient == nil {
		return errors.New("redis client not initialized - call InitRedis first")
	}
	return RClient.Set(context.TODO(), key, value, expiration).Err()
}

func Get(key string) (val string, err error) {
	if RClient == nil {
		return val, errors.New("redis client not initialized - call InitRedis first")
	}

	val, err = RClient.Get(context.TODO(), key).Result()
	if err == redis.Nil {
		return "", errors.New("Key does not exist")
	} else if err != nil {
		return "", err
	}
	return val, err
}

// TODO: Add Tests
func GetBytes(key string) (byts []byte, err error) {
	if RClient == nil {
		return byts, errors.New("redis client not initialized - call InitRedis first")
	}

	byts, err = RClient.Get(context.TODO(), key).Bytes()
	if err == redis.Nil {
		return byts, errors.New("Key does not exist")
	}

	return byts, err
}

// Return keys matching a pattern
func Scan(pattern string) (keys []string, err error) {
	if RClient == nil {
		return keys, errors.New("redis client not initialized - call InitRedis first")
	}

	var cursor uint64
	var errs []error
	for {
		var batchKeys []string
		batchKeys, cursor, err = RClient.Scan(context.TODO(), cursor, pattern, 15).Result()
		if err != nil {
			errs = append(errs, err)
		}
		keys = append(keys, batchKeys...)
		if cursor == 0 {
			break
		}
	}

	return
}

func Del(key string) error {
	if RClient == nil {
		return errors.New("redis client not initialized - call InitRedis first")
	}
	return RClient.Del(context.TODO(), key).Err()
}
