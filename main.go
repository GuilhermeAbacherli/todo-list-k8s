package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/GuilhermeAbacherli/todolistgo/service"
	"github.com/GuilhermeAbacherli/todolistgo/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	fmt.Println("\nGO todolist")
	fmt.Println("\nBem-vindo!")
	fmt.Println("O webservice funciona em paralelo com o terminal")
	utils.PressEnterKeyToContinue(reader)
}

func printMenu(reader *bufio.Reader) (stop bool) {
	fmt.Print("\nMenu\n\n")
	fmt.Println("0. Encerrar")
	fmt.Println("1. Listar TODOs")
	fmt.Println("2. Adicionar um TODO")
	fmt.Println("3. Completar um TODO")
	fmt.Println("4. Editar um TODO")
	fmt.Println("5. Excluir um TODO")
	fmt.Println("6. Excluir todos os TODOs")

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
		fmt.Println("\n3. Concluir um TODO")
		service.CompleteTodo(reader)
	case "4":
		clearTerminal()
		fmt.Println("\n4. Editar um TODO")
		service.EditTodo(reader)
	case "5":
		clearTerminal()
		fmt.Println("\n4. Excluir um TODO")
		service.RemoveTodo(reader)
	case "6":
		clearTerminal()
		fmt.Println("\n6. Excluir todos os TODOs")
		service.RemoveAllTodos(reader)
	default:
		fmt.Println("Escolha inválida, tente novamente")
	}
	return false
}

// GetClient returns a new mongodb client
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	currentClientConnection, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = currentClientConnection.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return currentClientConnection
}

func main() {

	clearTerminal()
	log.Println("Started")

	currentClientConnection := GetClient()
	err := currentClientConnection.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	dc := service.DatabaseConnection{
		Client: currentClientConnection,
	}

	// go func() {
	router := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})

	router.HandleFunc("/register", dc.Register).Methods("POST")
	router.HandleFunc("/login", dc.Login).Methods("POST")
	router.HandleFunc("/profile", dc.Profile).Methods("GET")

	router.HandleFunc("/todo", dc.GetManyTodos).Methods("GET")
	router.HandleFunc("/todo/done/{status}", dc.GetManyTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", dc.GetTodo).Methods("GET")
	router.HandleFunc("/todo", dc.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", dc.UpdateTodo).Methods("PATCH")
	router.HandleFunc("/todo/{id}", dc.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo", dc.DeleteAllTodos).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	// }()

	// reader := bufio.NewReader(os.Stdin)
	// printWelcome(reader)

	// for {
	// 	clearTerminal()
	// 	stop := printMenu(reader)
	// 	if stop {
	// 		break
	// 	}
	// }
}
