package validations

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func GetTodoValidation(c *gin.Context) {
	todoId, available := c.Params.Get("todoId")
	if available != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todoId is not passed",
		})
		c.Abort()
		return
	}

	if matched, _ := regexp.MatchString("\\b[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}\\b", todoId); matched != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todoId is not an UUID",
		})
		c.Abort()
		return
	}

	c.Next()
}

func CreateTodoValidation(c *gin.Context) (CreateTodoModel, error) {
	var todo CreateTodoModel
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation Error",
			"error":   err.Error(),
		})
		return CreateTodoModel{}, err
	}

	return todo, nil
}

type CreateTodoModel struct {
	ID    string `json:"id"`
	Title string `json:"title" binding:"min=3"`
}
