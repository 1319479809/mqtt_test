package auth

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/1319479809/mqtt_test/utils"
	"github.com/go-redis/redis/v8"
)

var (
	// tokenPool *redis.Pool
	// redisProtocol = "tcp"
	// redisAddr     = ":6379"
	// redisOptions  = []redis.DialOption{}

	rAuthClient  *redis.Client
	rTokenClient *redis.Client
	ctx          = context.Background()
)

func init() {
	initRedisToken()
	initRedisAuth()
	initRedisCache()
}

func initRedisToken() {
	redisAuthConf, err := utils.Cfg.GetSection("redis_token")
	if err != nil {
		panic("Redis 配置文件错误")
	}

	redisAddr := redisAuthConf.Key("addr").String()
	if redisAddr == "" {
		redisAddr = ":6379"
	}
	password := redisAuthConf.Key("password").String()

	db, _ := redisAuthConf.Key("db").Int()
	options := redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       db,
	}

	rTokenClient = redis.NewClient(&options)
	_, err = rTokenClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

}

func initRedisAuth() {
	redisAuthConf, err := utils.Cfg.GetSection("redis_auth")
	if err != nil {
		panic("Redis 配置文件错误")
	}

	redisAddr := redisAuthConf.Key("addr").String()
	if redisAddr == "" {
		redisAddr = ":6379"
	}
	password := redisAuthConf.Key("password").String()

	db, _ := redisAuthConf.Key("db").Int()
	options := redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       db,
	}

	rAuthClient = redis.NewClient(&options)
	_, err = rAuthClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

}

// EXPIRESIN 会话有效时间
const (
	EXPIRESIN = 24 * time.Hour
)

func SHA1Byte(input []byte) string {
	c := sha1.New()
	_, err := c.Write(input)
	if err != nil {
		return ""
	}
	bytes := c.Sum(nil)
	return hex.EncodeToString(bytes)
}

func SHA1(input string) string {
	c := sha1.New()
	_, err := c.Write([]byte(input))
	if err != nil {
		return ""
	}
	bytes := c.Sum(nil)
	return hex.EncodeToString(bytes)
}
