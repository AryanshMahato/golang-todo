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

	server.GET("/", controllers.GetTodosController)
	server.GET("/todo/:todoId", controllers.GetTodoController)
	server.DELETE("/todo/:todoId", controllers.DeleteTodoController)
	server.PUT("/todo/:todoId", controllers.UpdateTodoController)
	server.POST("/", controllers.CreateTodoController)

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

	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable",
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
