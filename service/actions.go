package service

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/GuilhermeAbacherli/todolistgo/entity"
	"github.com/GuilhermeAbacherli/todolistgo/utils"
)

var todoList []entity.Todo

// ShowTodo show a specif todo
func ShowTodo(todo *entity.Todo) {
	if todo != nil {
		fmt.Printf("\nTítulo: %s", todo.Title)
		fmt.Printf("\nDescrição: %s", todo.Description)
		if todo.Done {
			fmt.Println("\nConcluído: Sim")
		} else {
			fmt.Println("\nConcluído: Não")
		}
	}
}

// ShowTodoList print the todos in a list format
func ShowTodoList(index int, todo entity.Todo) {
	fmt.Printf("\n%d- ", (index + 1))
	fmt.Printf("Título: %s", todoList[index].Title)
	fmt.Printf("\nDescrição: %s", todoList[index].Description)
	if todoList[index].Done {
		fmt.Printf("\nConcluída: Sim")
	} else {
		fmt.Printf("\nConcluída: Não")
	}
	fmt.Print("\n----------------------")
}

// SearchTodo search for a specific todo based on the index or title
func SearchTodo(reader *bufio.Reader) (index int, todo *entity.Todo) {
	if len(todoList) < 1 {
		fmt.Print("\nNo momento não existem TODOs para serem exibidos.")

	} else {
		fmt.Println("\nDeseja buscar pelo código ou título?")
		choice := utils.Input(reader, "1. Código | 2. Título: ")

		switch choice {
		case "1":
			typedValue := utils.Input(reader, "Digite o código: ")
			index, _ := strconv.Atoi(typedValue)
			if index < 1 || index > len(todoList) {
				fmt.Printf("O código %d não existe.\n", index)
				if len(todoList) == 1 {
					fmt.Println("Só existe um TODO na lista.")
				} else {
					fmt.Printf("Digite um número entre 1 e %d.\n", len(todoList))
				}
			} else {
				index = index - 1
				return index, &todoList[index]
			}
		case "2":
			title := utils.Input(reader, "Digite o título: ")
			for index, todo := range todoList {
				if todo.Title == title {
					return index, &todo
				}
			}
			fmt.Println("Nenhum TODO encontrado com o título informado")
		default:
			fmt.Println("Opção inválida, tente novamente...")
		}
	}
	return -1, nil
}

// ListTodos Exibe todos os TODOs
func ListTodos(reader *bufio.Reader) {
	if len(todoList) < 1 {
		fmt.Print("\nNo momento não existem TODOs para serem exibidos.")
	} else {
		fmt.Print("\nDeseja exibir quais TODOs?")
		choice := utils.Input(reader, "\n1. Pendentes | 2. Concluídos | 3. Todos: ")
		switch choice {
		case "1":
			todosPending := 0
			for i := 0; i < len(todoList); i++ {
				if todoList[i].Done == false {
					todosPending++
					ShowTodoList(i, todoList[i])
				}
			}
			if todosPending == 0 {
				fmt.Print("Não há TODOs pendentes.")
			}
		case "2":
			todosCompleted := 0
			for i := 0; i < len(todoList); i++ {
				if todoList[i].Done == true {
					todosCompleted++
					ShowTodoList(i, todoList[i])
				}
			}
			if todosCompleted == 0 {
				fmt.Print("Não há TODOs concluídos.")
			}
		case "3":
			for i := 0; i < len(todoList); i++ {
				ShowTodoList(i, todoList[i])
			}
		default:
			fmt.Println("Opção inválida, tente novamente...")
		}
	}
	utils.PressEnterKeyToContinue(reader)
}

// AddTodo add a todo
func AddTodo(reader *bufio.Reader) {
	title := utils.Input(reader, "\n    Digite o título: ")
	description := utils.Input(reader, " Digite a descrição: ")
	todo := entity.NewTodoNotDone(title, description)
	todoList = append(todoList, todo)
	fmt.Println("------------------------------")
	fmt.Println("\nTodo adicionado:")
	ShowTodo(&todo)
	utils.PressEnterKeyToContinue(reader)
}

// CompleteTodo mark a TODO as done
func CompleteTodo(reader *bufio.Reader) {
	_, todo := SearchTodo(reader)
	if todo != nil {
		if todo.Done {
			fmt.Print("\nEste TODO já está concluído")
		} else {
			todo.Done = true
		}
		ShowTodo(todo)
		utils.PressEnterKeyToContinue(reader)
	}
}

// EditTodo edit some specific TODO
func EditTodo(reader *bufio.Reader) {
	_, todo := SearchTodo(reader)
	ShowTodo(todo)
	fmt.Println("------------------------------")
	fmt.Println("\nDigite os novos valores")
	todo.Title = utils.Input(reader, "Título: ")
	todo.Description = utils.Input(reader, "Descrição: ")
	for {
		done := utils.Input(reader, "Concluído (1. Sim | 2. Não): ")
		if done == "1" {
			todo.Done = true
			break
		}
		if done == "2" {
			todo.Done = false
			break
		}
	}
	ShowTodo(todo)
	utils.PressEnterKeyToContinue(reader)
}

// DeleteTodo delete some specific TODO
func DeleteTodo(reader *bufio.Reader) {
	index, todoDeleted := SearchTodo(reader)
	copy(todoList[index:], todoList[index+1:])
	todoList = todoList[:len(todoList)-1]
	fmt.Println("------------------------------")
	fmt.Println("\nTodo excluído:")
	ShowTodo(todoDeleted)
	utils.PressEnterKeyToContinue(reader)
}

// DeleteAllTodos delete all TODOs
func DeleteAllTodos(reader *bufio.Reader) {
	if len(todoList) < 1 {
		fmt.Println("\nNo momento não existem TODOs para serem excluídos.")
	} else {

		oldQuantity := len(todoList)
		todoList = nil
		if todoList == nil {
			if oldQuantity == 1 {
				fmt.Printf("\n%d TODO foi excluído", oldQuantity)
			} else {
				fmt.Printf("\nTodos os %d TODOs foram excluídos", oldQuantity)
			}
		}
	}
	utils.PressEnterKeyToContinue(reader)
}
