package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/GuilhermeAbacherli/todolistgo/service"
	"github.com/GuilhermeAbacherli/todolistgo/utils"
)

var clear map[string]func() // Create a map for storing clear functions

func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearTerminal() {
	valueFunction, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                                  // if we defined a clear func for that platform:
		valueFunction() // we execute it
	} else { // unsupported platform
		// panic("The platform is unsupported! The terminal can't be cleared...")
		panic("A plataforma não é suportada! O terminal não pode ser limpo...")
	}
}

func printWelcome(reader *bufio.Reader) {
	fmt.Println("\nGO TODO")
	fmt.Println("\nBem-vindo!")
	key := utils.Input(reader, "\nPressione qualquer tecla para continuar... ")
	if key != "" {
		return
	}
}

func printMenu(reader *bufio.Reader) (stop bool) {
	fmt.Print("Menu\n\n")
	fmt.Println("0. Encerrar")
	fmt.Println("1. Listar TODOs")
	fmt.Println("2. Adicionar um TODO")
	fmt.Println("3. Editar um TODO")
	fmt.Println("4. Excluir um TODO")
	fmt.Println("5. Excluir todos os TODOs")

	choice := utils.Input(reader, "\nDigite a opção desejada: ")

	switch choice {
	case "0":
		clearTerminal()
		fmt.Println("\n0. Encerrar")
		fmt.Println("\nTem certeza que deseja encerrar?")
		stop := utils.Input(reader, "\n1. Sim | 2. Cancelar: ")
		if stop == "1" {
			return true
		}
	case "1":
		clearTerminal()
		fmt.Println("\n1. Listar TODOs")
		service.ListTodos(reader)
	case "2":
		clearTerminal()
		fmt.Println("\n2. Adicionar um TODO")
		service.AddTodo(reader)
	case "3":
		clearTerminal()
		fmt.Println("\n3. Editar um TODO")
		service.EditTodo(reader)
	case "4":
		clearTerminal()
		fmt.Println("\n4. Excluir um TODO")
	case "5":
		clearTerminal()
		fmt.Println("\n5. Excluir todos os TODOs")
	default:
		fmt.Println("Escolha inválida, tente novamente")
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	clearTerminal()
	printWelcome(reader)

	for {
		clearTerminal()
		stop := printMenu(reader)
		if stop {
			break
		}
	}
}
