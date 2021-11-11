package redisdb2

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

func SETPX(conn *redis.Conn, key, value string, millisec uint) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	//SET 接口定义：SET key value PX milliseconds
	_, err = (*conn).Do("SET", key, value, "PX", millisec)
	return err
}

func TTL(conn *redis.Conn, key string) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	// TTL 接口定义：TTL key
	_, err = (*conn).Do("TTL", key)
	return err
}

func EXEC(conn *redis.Conn, cmdName string, params ...interface{}) (repl interface{}, err error) {
	if conn == nil {
		return nil, errors.New("connection is nil")
	}
	// TTL 接口定义：TTL key
	repl, err = (*conn).Do(cmdName, params...)
	return
}
