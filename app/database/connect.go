package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // _ allows us to import modules that are needed but not "used"
	"os"
	"thoughts/app/exceptions"
	"thoughts/app/helpers"
)

// Connect will generate an url with env variables and use that to connect to the database
// if all details are valid it will then return a connection
func Connect() *sql.DB {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		helpers.Envs["DB_USER"],
		helpers.Envs["DB_PASS"],
		helpers.Envs["DB_HOST"],
		helpers.Envs["DB_SCHEMA"])

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		exceptions.Error("Error while connecting to the database.", err)
		os.Exit(1)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		exceptions.Error("Error while pinging the database.", err)
		os.Exit(1)
	}

	return db
}