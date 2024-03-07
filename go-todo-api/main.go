package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:id`
	Item      string `json:item`
	Completed bool   `json:completed`
}

var todos = []todo{
	{ID: "1", Item: "cleaning room", Completed: false},
	{ID: "2", Item: "learning golang", Completed: false},
	{ID: "3", Item: "c++ project", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) //todos array data converted into json
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)

}

func toogleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodosById(id string) (*todo, error) {

	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {

	prot := ":4000"

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PUT("/todos/:id", toogleTodoStatus)

	fmt.Println("Server Listening....")
	err := router.Run(prot)
	if err != nil {
		log.Fatal(err)
	}
}

/*test apis

http://localhost:4000/todos  	get all todos
http://localhost:4000/todos		create todod
http://localhost:4000/todos/3	get todo by id
http://localhost:4000/todos/3	toggle completed satutus

*/
