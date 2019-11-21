package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/guilherme/todolist/entity"
)

var todoList []Todo

func printWelcome() {
	fmt.Println("\nGO TODO")
	fmt.Println("\nBem-vindo!")
	fmt.Println("\nPressione qualquer tecla para acessar o menu")
}

func printMenu() {
	fmt.Println("\nMenu")
	fmt.Println("1. Verificar lista de TODOs")
	fmt.Println("2. Adicionar um TODO")
	fmt.Println("3. Alterar um TODO")
	fmt.Println("4. Excluir um TODO")
	fmt.Println("5. Excluir todos os TODOs")
	fmt.Println("\nDigite a opção selecionada: ")
}

func input(reader *bufio.Reader, ask string) (text string) {
	fmt.Printf("%s:", ask)
	text, _ = reader.ReadString('\n')
	return
}

func addTodo() {
	reader := bufio.NewReader(os.Stdin)
	todo := entity.NewTodo()
	todo.Title = input(reader, "Digite o título")
	todo.Description = input(reader, "Digite a descrição")
	append(todoList, todo)
	//fmt.Printf("%v", todo)
}

func editTodo(reader *bufio.Reader) Todo {
	title := input(reader, "Digite o titulo do TODO que deseja alterar")
	oldTodo := todoList[oldTodo.Title]
	return newTodo
}

func main() {
	printMenu()

}
