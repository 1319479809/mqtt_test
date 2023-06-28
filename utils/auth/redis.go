package auth

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"time"

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
	initRedisCache()
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
