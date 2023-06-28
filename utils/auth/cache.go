package auth

import (
	"encoding/json"
	"log"
	"time"

	"github.com/1319479809/mqtt_test/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var (
	rCacheClient *redis.Client
)

const (
	CacheExpiresIn = 1 * time.Hour
)

func initRedisCache() {
	redisAuthConf, err := utils.Cfg.GetSection("redis_cache")
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

	log.Println("redis cache options: ", options.Addr, " Password:", options.Password, " DB:", options.DB)
	rCacheClient = redis.NewClient(&options)
	_, err = rCacheClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func GetCache(uid string) (string, error) {
	kvs, err := rCacheClient.Get(ctx, uid).Result()
	if err == redis.Nil {
		return kvs, err
	} else if err != nil {
		panic(err)
	}
	return kvs, err

}

func GetCacheData(uid string, data interface{}) (err error) {
	buf, err := GetCache(uid)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(buf), data)
	return err
}

func AddCache(data string, timeout ...time.Duration) (uid string, err error) {
	uid = uuid.New().String()
	err = SetCache(uid, data, timeout...)
	return
}

func SetCache(uid string, data string, timeout ...time.Duration) (err error) {
	if len(timeout) > 0 {
		err = rCacheClient.Set(ctx, uid, data, timeout[0]).Err()
		return
	}
	err = rCacheClient.Set(ctx, uid, data, CacheExpiresIn).Err()
	return
}

func SetCacheData(uid string, data interface{}) (err error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = SetCache(uid, string(buf))
	return
}

func DeleteCache(uid string) (err error) {
	err = rCacheClient.Del(ctx, uid).Err()
	return
}
