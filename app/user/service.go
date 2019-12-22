package user

import (
	"github.com/dogukanayd/go-rest/app/cache"
	"github.com/dogukanayd/go-rest/app/utils"
)

type ActionInterface interface {
	RequestCount() bool
}

type Cache struct {
	Cache *cache.Service
}

func New() ActionInterface {
	return &Cache{
		Cache: cache.New(utils.GetReusableRedisConnection()),
	}
}

func (c Cache) RequestCount() bool {
	c.Cache.SetKeyIdentifiers(123, "donovan").Incr(cache.UserRequestCount)

	return true
}
