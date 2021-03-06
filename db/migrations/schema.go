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
  				 login VARCHAR(150) NOT NULL,
  				 password VARCHAR(150) NOT NULL,
  				 email VARCHAR(255) NOT NULL,
  				 age int(3) NOT NULL);`,

		drop:   `DROP TABLE users`,
	},
	{
		tableName: "posts",
		create:`CREATE TABLE IF NOT EXISTS posts(
				id SERIAL NOT NULL PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				description TEXT NOT NULL,
				author_id INT NOT NULL
				);`,
		drop: `DROP TABLE posts`,
	},
}