package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sqlx.DB
}

func DbConnection(cfg *Config) (*DB, error){
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
		panic("Не поддерживаемый sql driver")
	}

	db, err := sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg.Driver, " connected")
	return &DB{db}, nil
}