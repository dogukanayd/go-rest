package utils

import (
	"github.com/dogukanayd/go-rest/app/config"
	"github.com/go-redis/redis"
	"log"
	"time"
)

// RedisInterface interface
type RedisInterface interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) string
	Incr(key string) int
	Ping() error
	Close()
}

var redisConnection = redis.NewClient(config.GetRedisOptions(10))

// Redis struct
type Redis struct {
	Client *redis.Client
}

// GetReusableRedisConnection returns connection initialized on app start
func GetReusableRedisConnection() RedisInterface {
	return &Redis{
		Client: redisConnection,
	}
}

// Increment the key
func (r Redis) Incr(key string) int {
	val := r.Client.Incr(key)

	return int(val.Val())
}

// Set key value with TTL
func (r Redis) Set(key string, value interface{}, ttl time.Duration) {
	r.Client.Set(key, value, ttl)
}

// Get value of a key
func (r Redis) Get(key string) string {
	res, err := r.Client.Get(key).Result()

	if err != nil {
		log.Println(err)

		return ""
	}

	return res
}

// Ping redis
func (r Redis) Ping() error {
	return r.Client.Ping().Err()
}

// Close connection
func (r Redis) Close() {
	err := r.Client.Close()

	if err != nil {
		log.Println(err)
	}
}
