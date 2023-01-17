package database

import (
	"xdu-health-card/database/mysql"
	"xdu-health-card/database/redis"
	"xdu-health-card/utils/flags"
)

func Setup() (err error) {
	_, err = mysql.Connect(flags.MysqlHost, flags.MysqlPort, flags.MysqlUser, flags.MysqlPasswd, flags.MysqlDB)
	if err != nil {
		return
	}

	_, err = redis.Connect(flags.RedisHost, flags.RedisPort, flags.RedisUser, flags.RedisPasswd, flags.RedisDb, flags.RedisPoolSize)
	if err != nil {
		return
	}
	return
}
