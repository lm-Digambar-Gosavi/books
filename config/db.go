package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {
	connectionString := "root:root@tcp(127.0.0.1:3306)/books_db?parseTime=true"

	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("database not reachable: %v", err)
	}
	log.Println("Database connected successfully")
	return nil
}
