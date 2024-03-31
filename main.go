package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TODOController interface {
	Create(c *gin.Context)
	Check(c *gin.Context)
	List(c *gin.Context)
}

type TODO struct {
	Todos []TODOItem `json:"todos"`
}

type TODOItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

func (c *TODO) Create(ctx *gin.Context) {

	var form TODOItem
	if err := ctx.BindJSON(&form); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Todos = append(c.Todos, form)
	ctx.JSON(http.StatusOK, gin.H{"Message": "Todo List create"})
}

func (c *TODO) Check(ctx *gin.Context) {

	type Form struct {
		ID string `json:"id,omitempty"`
	}

	form := Form{}

	if err := ctx.BindJSON(&form); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
	}

	found := false
	for i := range c.Todos {
		if c.Todos[i].ID == form.ID {
			c.Todos[i].IsCompleted = true
			found = true
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID for todo list"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Completed"})
}

func (c *TODO) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.Todos)
}

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	todo := &TODO{}
	server.POST("/create", todo.Create)
	server.GET("/list", todo.List)
	server.POST("/check", todo.Check)
	err := server.Run()
	if err != nil {
		return
	}
}
