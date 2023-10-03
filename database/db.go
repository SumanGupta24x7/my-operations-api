package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDbConnection() (*sqlx.DB, error) {
	username := "root"
	password := "password"
	hostname := "localhost"
	port := "3306"
	database_name := "OPERATIONS"

	db_connection_str := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + database_name + "?parseTime=true"

	db, err := sqlx.Open("mysql", db_connection_str)

	err = db.Ping()

	if err != nil {
		fmt.Println("Error occurred while connecting to database: ", err)
	}

	return db, err
}
