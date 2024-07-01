package dao

import (
	"api/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Implicit import. database/sql package will use it.
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DB_connection)
	if err != nil {
		fmt.Println("error at sql.Open")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error at ping")
		db.Close()
		return nil, err
	}

	fmt.Println("Connection opened.")

	return db, nil

}
