package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/GuilhermeAbacherli/todolistgo/dao"
	"github.com/GuilhermeAbacherli/todolistgo/entity"
)

func logRequest(request *http.Request) {
	log.Println(request.Method + " " + request.Host + request.RequestURI)
}

// DatabaseConnection maintains a single connection obj to call the methods
type DatabaseConnection struct {
	Client *mongo.Client
}

// GetAllTodos return all the TODOs from database
func (dc *DatabaseConnection) GetAllTodos(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	todolist := dao.SelectAllTodos(dc.Client, bson.M{})
	json.NewEncoder(writer).Encode(todolist)
}

// GetTodo return a specific TODO based on the ID
func (dc *DatabaseConnection) GetTodo(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	todo := dao.SelectOneTodo(dc.Client, bson.M{"ID": id})
	json.NewEncoder(writer).Encode(todo)
}

// CreateTodo creates a new TODO
func (dc *DatabaseConnection) CreateTodo(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	var todo entity.Todo
	_ = json.NewDecoder(request.Body).Decode(&todo)
	todo = entity.NewTodoWithOptions(todo.Title, todo.Description, todo.Done)
	insertedID := dao.InsertOneTodo(dc.Client, todo)
	json.NewEncoder(writer).Encode(insertedID)
}

// UpdateTodo update parcially a existing TODO
func (dc *DatabaseConnection) UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	var todo entity.Todo
	_ = json.NewDecoder(request.Body).Decode(&todo)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	modifiedCount := dao.UpdateOneTodo(dc.Client, todo, bson.M{"id": id})
	json.NewEncoder(writer).Encode(modifiedCount)
}

// DeleteTodo delete a specific TODO based on the ID
func (dc *DatabaseConnection) DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	deletedCount := dao.DeleteOneTodo(dc.Client, bson.M{"id": id})
	json.NewEncoder(writer).Encode(deletedCount)
}

// DeleteAllTodos delete all todos from database
func (dc *DatabaseConnection) DeleteAllTodos(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	deletedCount := dao.DeleteAllTodos(dc.Client)
	json.NewEncoder(writer).Encode(deletedCount)
}
