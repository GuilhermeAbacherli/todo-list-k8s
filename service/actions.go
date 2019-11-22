package service

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	"github.com/GuilhermeAbacherli/todolistgo/entity"
	"github.com/GuilhermeAbacherli/todolistgo/utils"
)

var todoList []entity.Todo

// ShowTodo show a specif todo
func ShowTodo(reader *bufio.Reader, todo *entity.Todo) {
	fmt.Printf("%s", todo.Title)
	key := utils.Input(reader, "\n\nPressione qualquer tecla para continuar... ")
	if key != "" {
		return
	}
}

// SearchTodo search for a specific todo based on the index or title
func SearchTodo(reader *bufio.Reader) *entity.Todo {
	fmt.Println("\nDeseja buscar pelo código ou título?")
	choice := utils.Input(reader, "1. Código | 2. Título: ")

	switch choice {
	case "1":
		typedValue := utils.Input(reader, "Digite o código: ")
		index, _ := strconv.Atoi(typedValue)
		if index < 1 || index > len(todoList) {
			fmt.Printf("Digite um número entre 0 e %d\n", len(todoList))
			fmt.Printf("O código %d não existe\n", index)
		} else {
			index = index - 1
			return &todoList[index]
		}

	case "2":
		title := utils.Input(reader, "Digite o título: ")
		for _, todo := range todoList {
			if todo.Title == title {
				return &todo
			}
		}
	default:
		fmt.Println("Opção inválida, tente novamente...")
	}

	time.Sleep(5 * time.Second)
	return nil
}

// ListTodos Exibe todos os TODOs
func ListTodos(reader *bufio.Reader) {
	if len(todoList) < 1 {
		fmt.Print("\nNo momento não existem TODOs para serem exibidos.")
	} else {
		for i := 0; i < len(todoList); i++ {
			fmt.Printf("\n%d - %s\n    %s\n    %t\n___________________________",
				(i + 1), todoList[i].Title, todoList[i].Description, todoList[i].Done)
		}
	}

	key := utils.Input(reader, "\n\nPressione qualquer tecla para continuar... ")
	if key != "" {
		return
	}
}

// AddTodo add a todo
func AddTodo(reader *bufio.Reader) {
	title := utils.Input(reader, "Digite o título: ")
	description := utils.Input(reader, "Digite a descrição: ")
	todo := entity.NewTodoNotDone(title, description)
	todoList = append(todoList, todo)
	fmt.Printf("Todo adicionado:\n%v", todo)
}

// EditTodo edit some specific TODO
func EditTodo(reader *bufio.Reader) {
	todo := SearchTodo(reader)
	ShowTodo(reader, todo)
}

// DeleteTodo delete some specific TODO
func DeleteTodo(reader *bufio.Reader) {

}

// DeleteAllTodos delete all TODOs
func DeleteAllTodos(reader *bufio.Reader) {
	todoList = nil
	if todoList == nil {
		fmt.Print("\n Todos os TODOs foram excluídos")
	}
	key := utils.Input(reader, "\n\nPressione qualquer tecla para continuar... ")
	if key != "" {
		return
	}
}
