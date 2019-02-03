package configs

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/errors"
	"github.com/Ayupov-Ayaz/go-web/ui"
	"github.com/viper"
	"os"
	"strconv"
)

func GetUiConfig() *ui.Config{
	return &ui.Config{
		AssetsPrefix: "/assets/",
		AssetsPath: "./assets/",
	}
}

func GetDbConfig() *db.Config{
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
	fmt.Println(dbParams)
	port, err  := strconv.Atoi(dbParams["port"])
	if err != nil {
		errors.PrintSystemErr(fmt.Sprintf("При попытке получить port возникла ошибка: %s", err.Error()))
	}
	return &db.Config{
		Driver: dbParams["driver"],
		Host: dbParams["host"],
		Database: dbParams["database"],
		User: dbParams["user"],
		Password: dbParams["password"],
		Port: port,
	}
}