package flags

var (
	MysqlHost   string
	MysqlUser   string
	MysqlPasswd string
	MysqlDB     string
	MysqlPort   int

	RedisHost     string
	RedisPort     int
	RedisDb       int
	RedisUser     string
	RedisPasswd   string
	RedisPoolSize int

	AppID      string
	Secret     string
	HmacSecret string
	MapApiKey  string
)
