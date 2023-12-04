package main

import (
	"fmt"
	"go-comms/service/pkg/handlers"
	"go-comms/service/pkg/reader"
	"go-comms/service/pkg/store"
	"net/http"
	"os"
	"strings"
)

const (
	Port = 8080
	Exit = "exit"
)

func main() {

	go readInputFromConsole()
	http.HandleFunc("/messages", handlers.GetMessagesHandler)

	fmt.Printf("Starting your service...\nServer is listening on :%d...\n", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}

func readInputFromConsole() {
	cr := reader.ConsoleReader{}
	for {
		input := cr.Read()

		store.AddInput(input)
		if input != "" && strings.ToLower(input) == Exit {
			fmt.Println("Termiating your service...")
			os.Exit(0)
		}
	}
}
