package migrations

import (
	"fmt"
	"go-web/db"
)

func Migration(command string) {
	if command == "--migrate" {
		 err := startMigration()
		if err != nil {
			fmt.Println(err.Error())
		}
		 fmt.Println("Migration done")
	} else if command == "--rollback" {
		err := rollBack()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Rollback done")
	} else {
		fmt.Printf("Неизвестная команда \"%s\"\n", command)
	}
}

func startMigration()  error {
	db, err := getDbConnection()

	if err != nil {
		return err
	}

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

func rollBack() error {
	db, err := getDbConnection()
	if err != nil {
		return err
	}

	for _, schema := range allSchemas {
		_, err := db.Query(schema.drop)
		if err != nil {
			fmt.Println("Не удалось удалить таблицу:", schema.tableName, "\n")
			return err
		} else {
			fmt.Printf("Таблица \"%s\" - успешно удалена \n", schema.tableName)
		}
	}
	return nil
}

func getDbConnection() (*db.DB, error) {
	dbConfig := db.GetConfigs()
	return db.InitDB(dbConfig)
}
