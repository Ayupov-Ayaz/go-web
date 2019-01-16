package migrations

import (
	"fmt"
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

/**
 	Создание всех таблиц
 */
func createTables(db *db.DB) error {
	for _, schema := range allSchemas {
		_, err := db.Query(schema.create)
		if err != nil {
			fmt.Println("Не удалось создать таблицу:", schema.tableName, "\n", err.Error())
		} else {
			fmt.Printf("Таблица \"%s\" - успешно создана \n", schema.tableName)
		}
	}
	return nil
}

/**
	Удаление всех таблиц
 */
func dropTables(db *db.DB) error {
	for _, schema := range allSchemas {
		_, err := db.Query(schema.drop)
		if err != nil {
			fmt.Println("Не удалось удалить таблицу:", schema.tableName, "\n", err.Error())
		} else {
			fmt.Printf("Таблица \"%s\" - успешно удалена \n", schema.tableName)
		}
	}
	return nil
}