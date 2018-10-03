package rmq

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisWrapper struct {
	rawClient *redis.Client
}

func (wrapper RedisWrapper) Set(key string, value string, expiration time.Duration) error {
	return wrapper.rawClient.Set(key, value, expiration).Err()
}

func (wrapper RedisWrapper) Del(key string) (int, error) {
	n, err := wrapper.rawClient.Del(key).Result()
	return int(n), err
}

func (wrapper RedisWrapper) TTL(key string) (ttl time.Duration, err error) {
	return wrapper.rawClient.TTL(key).Result()
}

func (wrapper RedisWrapper) LPush(key, value string) (error) {
	return wrapper.rawClient.LPush(key, value).Err()
}

func (wrapper RedisWrapper) LLen(key string) (int, error) {
	n, err := wrapper.rawClient.LLen(key).Result()
	return int(n), err
}

func (wrapper RedisWrapper) LRem(key string, count int, value string) (int, error) {
	n, err := wrapper.rawClient.LRem(key, int64(count), value).Result()
	return int(n), err
}

func (wrapper RedisWrapper) LTrim(key string, start, stop int) error {
	return wrapper.rawClient.LTrim(key, int64(start), int64(stop)).Err()
}

func (wrapper RedisWrapper) RPopLPush(source, destination string) (value string, err error) {
	return wrapper.rawClient.RPopLPush(source, destination).Result()
}

func (wrapper RedisWrapper) SAdd(key, value string) (error) {
	return wrapper.rawClient.SAdd(key, value).Err()
}

func (wrapper RedisWrapper) SMembers(key string) (members []string, err error) {
	return wrapper.rawClient.SMembers(key).Result()
}

func (wrapper RedisWrapper) SRem(key, value string) (affected int, err error) {
	n, err := wrapper.rawClient.SRem(key, value).Result()
	return int(n), err
}

func (wrapper RedisWrapper) FlushDb() {
	wrapper.rawClient.FlushDb()
}
