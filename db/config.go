package db

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/errors"
	"github.com/viper"
	"os"
	"strconv"
)

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

func GetDbConfig() *Config{
	viper.SetConfigType("yaml")
	fileName := ".db.yaml"
	in, err := os.Open(fileName)
	if err != nil {
		errors.PrintSystemErr(fmt.Sprintf("Не удалось прочитать файл конфигурации db \n %s", err.Error()))
	}
	if err := viper.ReadConfig(in); err != nil {
		errors.PrintSystemErr(fmt.Sprintf("При чтении файла конфигурации db произашла ошибка \n %s",
			err.Error()))
	}
	dbParams := viper.GetStringMapString("db")
	port, err  := strconv.Atoi(dbParams["port"])
	if err != nil {
		errors.PrintSystemErr(fmt.Sprintf("При попытке получить port возникла ошибка: %s", err.Error()))
	}
	return &Config{
		Driver: dbParams["driver"],
		Host: dbParams["host"],
		Database: dbParams["database"],
		User: dbParams["user"],
		Password: dbParams["password"],
		Port: port,
	}
}