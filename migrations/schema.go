package migrations

type Schema struct {
	tableName string
	create string
	drop string
}

var allSchemas = []Schema{
	// table users
	{
		tableName: "users",
		create: `CREATE TABLE IF NOT EXISTS users (
  				 id SERIAL NOT NULL PRIMARY KEY,
  				 first_name VARCHAR(150) NOT NULL,
  				 last_name VARCHAR(150) NOT NULL,
  				 email VARCHAR(255) NOT NULL,
  				 age int(3) NOT NULL);`,

		drop:   `DROP TABLE users`,
	},
	//
}