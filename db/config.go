package db

const (
	DriverMSSQL  = "mssql"
	DriverMySQL  = "mysql"
)

type Config struct {
	Driver string
	User string
	Password string
	Host string
	Database string
	Port int
}

//TODO: getenv
func GetConfigs() *Config{
	return &Config{
		Driver: "mysql",
		User: "tommy",
		Password: "43",
		Host: "127.0.0.1",
		Database: "go_lang",
		Port: 3306,
	}
}