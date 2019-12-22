package cache

type RedisKey string

const (
	UserRequestCount RedisKey = "userName:{userName}:userID:{userID}:user-request-count"
)