package service

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/GuilhermeAbacherli/todolistgo/entity"
	"github.com/GuilhermeAbacherli/todolistgo/utils"
)

// ShowTodo show a specif todo
func ShowTodo(todo *entity.Todo) {
	if todo != nil {
		fmt.Printf("\nTítulo: '%s'", todo.Title)
		fmt.Printf("\nDescrição: '%s'", todo.Description)
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
	fmt.Printf("Título: %s", entity.TodoList[index].Title)
	fmt.Printf("\nDescrição: %s", entity.TodoList[index].Description)
	if entity.TodoList[index].Done {
		fmt.Printf("\nConcluída: Sim")
	} else {
		fmt.Printf("\nConcluída: Não")
	}
	fmt.Print("\n----------------------")
}

// SearchTodo search for a specific todo based on the index or title
func SearchTodo(reader *bufio.Reader) (index int, todo *entity.Todo) {
	if len(entity.TodoList) < 1 {
		fmt.Print("\nNo momento não existem TODOs para serem exibidos.")

	} else {
		fmt.Println("\nDeseja buscar pelo código ou título?")
		choice := utils.Input(reader, "1. Código | 2. Título: ")

		switch choice {
		case "1":
			typedValue := utils.Input(reader, "Digite o código: ")
			index, _ := strconv.Atoi(typedValue)
			if index < 1 || index > len(entity.TodoList) {
				fmt.Printf("O código %d não existe.\n", index)
				if len(entity.TodoList) == 1 {
					fmt.Println("Só existe um TODO na lista.")
				} else {
					fmt.Printf("Digite um número entre 1 e %d.\n", len(entity.TodoList))
				}
			} else {
				index = index - 1
				return index, &entity.TodoList[index]
			}
		case "2":
			title := utils.Input(reader, "Digite o título: ")
			for index, todo := range entity.TodoList {
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
	if len(entity.TodoList) < 1 {
		fmt.Print("\nNo momento não existem TODOs para serem exibidos.")
	} else {
		fmt.Print("\nDeseja exibir quais TODOs?")
		choice := utils.Input(reader, "\n1. Pendentes | 2. Concluídos | 3. Todos: ")
		switch choice {
		case "1":
			todosPending := 0
			for i := 0; i < len(entity.TodoList); i++ {
				if entity.TodoList[i].Done == false {
					todosPending++
					ShowTodoList(i, entity.TodoList[i])
				}
			}
			if todosPending == 0 {
				fmt.Print("\nNão há TODOs pendentes.")
			}
		case "2":
			todosCompleted := 0
			for i := 0; i < len(entity.TodoList); i++ {
				if entity.TodoList[i].Done == true {
					todosCompleted++
					ShowTodoList(i, entity.TodoList[i])
				}
			}
			if todosCompleted == 0 {
				fmt.Print("\nNão há TODOs concluídos.")
			}
		case "3":
			for i := 0; i < len(entity.TodoList); i++ {
				ShowTodoList(i, entity.TodoList[i])
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
	entity.TodoList = append(entity.TodoList, todo)
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
	if todo != nil {
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
	}
	utils.PressEnterKeyToContinue(reader)
}

// RemoveTodo delete some specific TODO
func RemoveTodo(reader *bufio.Reader) {
	index, todoToDelete := SearchTodo(reader)
	if todoToDelete != nil {
		copy(entity.TodoList[index:], entity.TodoList[index+1:])
		entity.TodoList = entity.TodoList[:len(entity.TodoList)-1]
		fmt.Println("------------------------------")
		fmt.Println("\nTodo excluído:")
		ShowTodo(todoToDelete)
	}
	utils.PressEnterKeyToContinue(reader)
}

// RemoveAllTodos delete all TODOs
func RemoveAllTodos(reader *bufio.Reader) {
	if len(entity.TodoList) < 1 {
		fmt.Println("\nNo momento não existem TODOs para serem excluídos.")
	} else {
		oldQuantity := len(entity.TodoList)
		entity.TodoList = nil
		if entity.TodoList == nil {
			if oldQuantity == 1 {
				fmt.Printf("\n%d TODO foi excluído", oldQuantity)
			} else {
				fmt.Printf("\nTodos os %d TODOs foram excluídos", oldQuantity)
			}
		}
	}
	utils.PressEnterKeyToContinue(reader)
}
