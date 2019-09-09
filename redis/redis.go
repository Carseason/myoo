package redis

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool  *redis.Pool
	RedisError error
)

func New(addr, pass string) {
	redisPool = &redis.Pool{
		MaxIdle:     30,
		MaxActive:   200,
		IdleTimeout: 30 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", addr, redis.DialPassword(pass))
		},
	}
	RedisPing()
}

func RedisPing() {
	conn := redisPool.Get()
	defer conn.Close()
	if _, err := conn.Do("PING"); err != nil {
		RedisError = err
		log.Println("Redis :", err)
	}
}
func Put(key string, value interface{}, redisTime int64) error {
	conn := redisPool.Get()
	defer conn.Close()
	var err error
	if redisTime != 0 {
		_, err = conn.Do("SET", key, value, "EX", redisTime) //, redisTime*time.Second
	} else {
		_, err = conn.Do("SET", key, value) //, redisTime*time.Second
	}
	return err
}
func PutStruct(key string, value interface{}, redisTime int64) error {
	conn := redisPool.Get()
	defer conn.Close()
	if jsonData, err := json.Marshal(value); err == nil {
		return Put(key, jsonData, redisTime)
	} else {
		return err
	}
}

func Get(name string) []byte {
	conn := redisPool.Get()
	defer conn.Close()
	if value, err := conn.Do("get", name); err != nil || value == nil {
		return []byte(nil)
	} else {
		return value.([]byte)
	}
}

func GetInt(key string) int {
	conn := redisPool.Get()
	defer conn.Close()
	if value, err := redis.Int(conn.Do("get", key)); err == nil {
		return value
	}
	return 0
}
func GetInt64(key string) int64 {
	conn := redisPool.Get()
	defer conn.Close()
	if value, err := redis.Int64(conn.Do("get", key)); err == nil {
		return value
	}
	return 0

}
func GetString(key string) string {
	conn := redisPool.Get()
	defer conn.Close()
	if value, err := redis.String(conn.Do("get", key)); err == nil {
		return value
	}
	return ""
}
func GetBool(key string) bool {
	conn := redisPool.Get()
	defer conn.Close()
	if value, err := redis.Bool(conn.Do("get", key)); err == nil {
		return value
	}
	return false

}
func GetStruct(key string, str interface{}) error {
	conn := redisPool.Get()
	defer conn.Close()
	if value, err := conn.Do("get", key); err == nil {
		return json.Unmarshal(value.([]byte), str)
	} else {
		return err
	}

}
func Is(name string) bool {
	if RedisError != nil {

		return false
	}

	conn := redisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", name))
	if err != nil {
		return false
	}
	return exists
}

func Del(name string) error {
	conn := redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("del", name)
	return err
}
func DelAll(key string) error {
	if key == "" {
		return nil
	}
	conn := redisPool.Get()
	defer conn.Close()
	value, err := redis.Strings(conn.Do("keys", key))
	if err != nil {
		return err
	}
	for v := range value {
		conn.Do("del", value[v])

	}
	return nil
}

/*************自增***********/
func Incr(key string) error {
	conn := redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("INCR", key)
	return err
}
