package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
	"time"
	"xdu-health-card/model/base"
)

var rdb *redis.Client

func Connect(IP string, port int, username, password string, db int, poolSize int) (Rdb *redis.Client, err error) {
	dbAddr := fmt.Sprintf("%s:%d", IP, port)
	rdb = redis.NewClient(&redis.Options{
		Username:     username,
		Addr:         dbAddr,
		Password:     password,
		DB:           db,
		PoolSize:     poolSize,
		MaxRetries:   3,
		MinIdleConns: poolSize / 2,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		log.Errorf("Connect to redis '%s'@'%s:%d', DB:%d failed", username, IP, port, db)
		return
	}
	log.Infof("Successfully connected to redis '%s'@'%s:%d', DB:%d", username, IP, port, db)
	Rdb = rdb
	return
}

func Subscribe(rdb *redis.Client, channels ...string) (ch <-chan *redis.Message, err error) {
	if len(channels) == 0 {
		return
	}
	pubsub := rdb.Subscribe(channels...)
	_, err = pubsub.Receive()
	if err != nil {
		log.Fatal(err)
	}
	ch = pubsub.Channel()
	return
}

func Get(rdb *redis.Client, key string) (val string, err error) {
	val, err = rdb.Get(key).Result()
	return
}

func Set(rdb *redis.Client, key, value string) (val string, err error) {
	val, err = rdb.Set(key, value, 0).Result()
	return
}

func Del(rdb *redis.Client, key string) (val int64, err error) {
	val, err = rdb.Del(key).Result()
	return
}

func SetNX(rdb *redis.Client, key, value string, expiration time.Duration) (succeed bool, err error) {
	succeed, err = rdb.SetNX(key, value, expiration).Result()
	return
}

func lock(key, value string) (succeed bool, err error) {
	return SetNX(rdb, key, value, 10*time.Second)
}

func Lock(key string) (retCode *base.RetCode, err error) {
	succeed, err := lock(key, "")
	if err != nil {
		return base.SetRedisLockFailed, err
	}
	if !succeed {
		return base.LockExist, nil
	}
	return base.SUCCESS, nil
}

func Unlock(key string) (val int64, err error) {
	return Del(rdb, key)
}
