package database

import (
	"xdu-health-card/database/mysql"
	"xdu-health-card/utils/flags"
)

func Setup() (err error) {
	_, err = mysql.Connect(flags.MysqlHost, flags.MysqlPort, flags.MysqlUser, flags.MysqlPasswd, flags.MysqlDB)
	if err != nil {
		return
	}
	return
}
