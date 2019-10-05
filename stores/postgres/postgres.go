package postgres

import "github.com/go-pg/pg"

// CreateDbConnection crea una conexion a la base de datos postgres
func CreateDbConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "root",
		Password: "password",
		Database: "gochat",
	})

	return db
}
