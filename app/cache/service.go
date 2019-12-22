package cache

import (
	"fmt"
	"github.com/dogukanayd/go-rest/app/utils"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	redis    utils.RedisInterface
	userName string
	userID   string
	key      string
}

// New cache service instance
func New(redis utils.RedisInterface) *Service {
	return &Service{
		redis: redis,
	}
}

// SetIdentifiers for the key
func (s *Service) SetKeyIdentifiers(userID int, userName string) *Service {
	s.userID = strconv.Itoa(userID)
	s.userName = userName

	return s
}

// Cache data
func (s *Service) Cache(key RedisKey, data interface{}, ttl time.Duration) *Service {
	s.redis.Set(s.getKey(key), data, ttl)

	return s
}

// Increment data
func (s *Service) Incr(key RedisKey) int {
	return s.redis.Incr(s.getKey(key))
}

func (s *Service) getKey(key RedisKey) string {
	return strings.NewReplacer("{userName}", s.userName, "{userID}", s.userID).Replace(fmt.Sprint(key))
}
