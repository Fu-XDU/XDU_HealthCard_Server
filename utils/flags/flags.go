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
		Name:        "mysqlHost",
		Usage:       "Mysql host",
		Value:       "127.0.0.1",
		EnvVar:      "MYSQL_HOST",
		Destination: &MysqlHost,
	}
	MysqlPortFlag = cli.IntFlag{
		Name:        "mysqlPort",
		Usage:       "Mysql port",
		Value:       3306,
		EnvVar:      "MYSQL_PORT",
		Destination: &MysqlPort,
	}
	MysqlUserFlag = cli.StringFlag{
		Name:        "mysqlUser",
		Usage:       "Mysql user",
		Value:       "root",
		EnvVar:      "MYSQL_USER",
		Destination: &MysqlUser,
	}
	MysqlPasswdFlag = cli.StringFlag{
		Name:        "mysqlPasswd",
		Usage:       "Mysql password",
		Value:       "123456",
		EnvVar:      "MYSQL_PASSWD",
		Destination: &MysqlPasswd,
	}
	MysqlDBFlag = cli.StringFlag{
		Name:        "mysqlDB",
		Usage:       "Mysql database",
		Value:       "xdu_health_card",
		EnvVar:      "MYSQL_DB",
		Destination: &MysqlDB,
	}

	AppidFlag = cli.StringFlag{
		Name:        "appid",
		Usage:       "Appid",
		Value:       "",
		EnvVar:      "APPID",
		Destination: &AppID,
	}
	SecretFlag = cli.StringFlag{
		Name:        "secret",
		Usage:       "App secret",
		Value:       "",
		EnvVar:      "SECRET",
		Destination: &Secret,
	}
	HmacSecretFlag = cli.StringFlag{
		Name:        "hmacSecret",
		Usage:       "Hmac secret, a random string",
		Value:       "",
		EnvVar:      "HMAC_SECRET",
		Destination: &HmacSecret,
	}
	MapApiKeyFlag = cli.StringFlag{
		Name:        "mapApiKey",
		Usage:       "Map api key",
		Value:       "",
		EnvVar:      "MAP_API_KEY",
		Destination: &MapApiKey,
	}
)
