package migrations

import (
	"fmt"
	"go-web/db"
)

/**
 	Создание всех таблиц
 */
func CreateTables(db *db.DB) error {
	for _, schema := range allSchemas {
		_, err := db.Query(schema.create)
		if err != nil {
			fmt.Println("Не удалось создать таблицу:", schema.tableName, "\n", err.Error())
		}
	}
	return nil
}

/**
	Удаление всех таблиц
 */
func DropTables(db *db.DB) error {
	for _, schema := range allSchemas {
		_, err := db.Query(schema.drop)
		if err != nil {
			fmt.Println("Не удалось удалить таблицу:", schema.tableName, "\n", err.Error())
		}
	}
	return nil
}