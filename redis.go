package redisdb2

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	serverTags map[string]*redis.Pool = make(map[string]*redis.Pool)
)

func New(serverTag, server, password string, dbnum int) {
	redisPool := &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: 240 * time.Second,
		MaxActive:   1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(dbnum))
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	serverTags[serverTag] = redisPool
}

func Destroy() {
	for k := range serverTags {
		delete(serverTags, k)
	}
}
