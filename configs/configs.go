package configs

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/ui"
)

func GetUiConfig() *ui.Config{
	return &ui.Config{
		AssetsPrefix: "/assets/",
		AssetsPath: "./assets/"}
}

func GetDbConfig() *db.Config{
	return &db.Config{
		Driver: "mysql",
		Host: "127.0.0.1",
		Database: "go_lang",
		User: "tommy",
		Password: "43",
		Port: 3306}
}