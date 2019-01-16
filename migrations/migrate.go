package migrations


import (
	"go-web/db"
)

func StartMigration() {
	connect, err := getDbConnection()
	if err != nil {
		panic(err)
	}
	createTables(connect)
}

func RollBack() {
	connect, err := getDbConnection()
	if err != nil {
		panic(err)
	}
	dropTables(connect)
}

func getDbConnection() (*db.DB, error) {
	dbConfig := db.GetConfigs()
	return db.InitDB(dbConfig)
}
