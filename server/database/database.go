package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "protokoll",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
