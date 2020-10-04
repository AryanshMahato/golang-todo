package main

import (
	"GoLang/controllers"
	"GoLang/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	server := gin.Default()

	connectDatabase()

	server.GET("/health", controllers.HealthController)

	err := server.Run()
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}

func connectDatabase() {
	var (
		user     = os.Getenv("user")
		dbname   = os.Getenv("dbname")
		password = os.Getenv("password")
	)

	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disabled",
		user,
		dbname,
		password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Database Connection failed", err)
		return
	}

	model.DB = db
}
