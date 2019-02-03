package db

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func InitDB(cfg *Config) (*DB, error){
	var dsn string

	switch cfg.Driver {
	case DriverMSSQL:
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
	case DriverMySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
	default:
		errors.PrintSystemErr(fmt.Sprintf("%s - не поддерживаемый sql driver", cfg.Driver))
	}

	db, err := sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg.Driver, " connected")
	return &DB{db}, nil
}