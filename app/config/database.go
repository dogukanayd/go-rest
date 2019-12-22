package config

import (
	"github.com/go-redis/redis"
	"os"
)

var connections = map[string]string{
	"mysql": os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@" + os.Getenv("MYSQL_HOST"),
}

// GetDBDSN returns connection dsn
func GetDBDSN(dbName string) string {
	return connections["mysql"] + "/" + dbName
}

func GetRedisOptions(poolSize int) *redis.Options {
	return &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
		PoolSize: poolSize,
	}
}
