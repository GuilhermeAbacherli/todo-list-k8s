package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/GuilhermeAbacherli/todolistgo/entity"
)

// GetAllTodos return all the TODOs from database
func GetAllTodos(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(entity.TodoList)
}

// GetTodo return a specific TODO based on the ID
func GetTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	for _, todo := range entity.TodoList {
		if todo.ID == id {
			json.NewEncoder(writer).Encode(todo)
			return
		}
	}
	return
}

// CreateTodo creates a new TODO
func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var todo entity.Todo
	_ = json.NewDecoder(request.Body).Decode(&todo)
	todo = entity.NewTodoWithOptions(todo.Title, todo.Description, todo.Done)
	entity.TodoList = append(entity.TodoList, todo)
	json.NewEncoder(writer).Encode(todo)
}

// UpdateTodo update parcially a existing TODO
func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	for index, todo := range entity.TodoList {
		if todo.ID == id {
			_ = json.NewDecoder(request.Body).Decode(&todo)
			entity.TodoList[index] = todo
			json.NewEncoder(writer).Encode(todo)
			break
		}
	}
}

// DeleteTodo delete a specific TODO based on the ID
func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	for index, todo := range entity.TodoList {
		if todo.ID == id {
			entity.TodoList = append(entity.TodoList[:index], entity.TodoList[index+1:]...)
			json.NewEncoder(writer).Encode(todo)
			break
		}
	}
}

// DeleteAllTodos delete all todos from database
func DeleteAllTodos(writer http.ResponseWriter, request *http.Request) {
	oldQuantity := len(entity.TodoList)
	if oldQuantity == 0 {
		json.NewEncoder(writer).Encode(-1)
	} else {
		entity.TodoList = nil
		json.NewEncoder(writer).Encode(oldQuantity)
	}
}
