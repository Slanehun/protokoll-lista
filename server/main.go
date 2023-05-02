package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/Slanehun/protokoll-lista/database"
	"github.com/Slanehun/protokoll-lista/rest"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to MySQL database!")

	// Register rest endpoints
	r := gin.Default()
	rest.RegisterEndpoints(r, db)

	// Run the app
	log.Fatal(r.Run(":8080"))
}
