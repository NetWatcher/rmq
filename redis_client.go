package rmq

import "time"

type RedisClient interface {
	// simple keys
	Set(key string, value string, expiration time.Duration) error
	Del(key string) (affected int, err error)      // default affected: 0
	TTL(key string) (ttl time.Duration, err error) // default ttl: 0

	// lists
	LPush(key, value string) error
	LLen(key string) (affected int, err error)
	LRem(key string, count int, value string) (affected int, err error)
	LTrim(key string, start, stop int) error
	RPopLPush(source, destination string) (value string, err error)

	// sets
	SAdd(key, value string) error
	SMembers(key string) (members []string, err error)         // default members: []string{}
	SRem(key, value string) (affected int, err error) // default affected: 0

	// special
	FlushDb()
}
