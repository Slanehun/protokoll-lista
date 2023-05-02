package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

func Connect() (*sql.DB, error) {
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
