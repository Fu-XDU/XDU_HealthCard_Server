package flags

import "gopkg.in/urfave/cli.v1"

var (
	PortFlag = cli.StringFlag{
		Name:   "port, p",
		Usage:  "Server port",
		Value:  "1423",
		EnvVar: "SERVER_PORT",
	}

	MysqlHostFlag = cli.StringFlag{
		Name:        "MysqlHost",
		Usage:       "MysqlHost",
		Value:       "127.0.0.1",
		EnvVar:      "MYSQL_HOST",
		Destination: &MysqlHost,
	}
	MysqlPortFlag = cli.IntFlag{
		Name:        "MysqlPort",
		Usage:       "MysqlPort",
		Value:       3306,
		EnvVar:      "MYSQL_PORT",
		Destination: &MysqlPort,
	}
	MysqlUserFlag = cli.StringFlag{
		Name:        "MysqlUser",
		Usage:       "MysqlUser",
		Value:       "root",
		EnvVar:      "MYSQL_USER",
		Destination: &MysqlUser,
	}
	MysqlPasswdFlag = cli.StringFlag{
		Name:        "MysqlPasswd",
		Usage:       "MysqlPasswd",
		Value:       "123456",
		EnvVar:      "MYSQL_PASSWD",
		Destination: &MysqlPasswd,
	}
	MysqlDBFlag = cli.StringFlag{
		Name:        "MysqlDB",
		Usage:       "mysqlDatabase",
		Value:       "xdu_health_card",
		EnvVar:      "MYSQL_DB",
		Destination: &MysqlDB,
	}

	AppidFlag = cli.StringFlag{
		Name:        "Appid",
		Usage:       "Appid",
		Value:       "",
		EnvVar:      "APPID",
		Destination: &AppID,
	}
	SecretFlag = cli.StringFlag{
		Name:        "Secret",
		Usage:       "Secret",
		Value:       "",
		EnvVar:      "SECRET",
		Destination: &Secret,
	}
)
