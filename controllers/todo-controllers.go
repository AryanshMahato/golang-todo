package controllers

import (
	"GoLang/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"alive": true,
	})
}

func GetTodosController(ctx *gin.Context) {
	rows, err := model.DB.Query("SELECT * FROM TODO")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot fetch todos",
			"todos":   rows,
		})
		fmt.Println(err)
		return
	}

	var todos []model.Todo

	defer rows.Close()
	for rows.Next() {
		var id string
		var title string
		err = rows.Scan(&id, &title)
		todos = append(todos, model.Todo{
			ID:    id,
			Title: title,
		})
	}

	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func GetTodoController(ctx *gin.Context) {
	todoId, err := ctx.Params.Get("todoId")
	if err == false {
		ctx.JSON(400, gin.H{
			"message": "todoId is not passed",
		})
		return
	}

	row := model.DB.QueryRow("SELECT * FROM Todo WHERE ID=$1", todoId)

	var todo model.Todo
	scanError := row.Scan(&todo.ID, &todo.Title)
	if scanError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": scanError.Error(),
		})
		fmt.Println(scanError)
		return
	}

	ctx.JSON(200, gin.H{
		"todo": todo,
	})
}

func CreateTodoController(ctx *gin.Context) {
	body := model.Todo{}
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot create todo",
		})
		return
	}

	row := model.DB.QueryRow("INSERT INTO TODO(Title) VALUES ($1) RETURNING ID", body.Title)

	var todoId string
	err = row.Scan(&todoId)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Todo Created",
		"todoId":  todoId,
	})
}

func DeleteTodoController(ctx *gin.Context) {
	todoId, available := ctx.Params.Get("todoId")
	if available != true {
		ctx.JSON(400, gin.H{
			"message": "todoId is not passed",
		})
		return
	}

	_ = model.DB.QueryRow("DELETE FROM Todo WHERE ID=$1", todoId)

	ctx.JSON(200, gin.H{
		"message": "todo deleted",
		"todoId":  todoId,
	})
}

func UpdateTodoController(ctx *gin.Context) {
	todoId, available := ctx.Params.Get("todoId")
	if available != true {
		ctx.JSON(400, gin.H{
			"message": "todoId is not passed",
		})
		return
	}

	body := model.Todo{}
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot update todo",
		})
		return
	}

	row := model.DB.QueryRow("UPDATE Todo SET Title=$1 WHERE ID=$2 RETURNING ID", body.Title, todoId)

	err = row.Scan(&todoId)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "todo update failed",
			"error":   err.Error(),
		},
		)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo updated",
	})
}
