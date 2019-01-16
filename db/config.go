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

