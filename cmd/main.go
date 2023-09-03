package main

import (
	"awesomeProject1/cmd/service"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Key-Value Store! Commands:")
	fmt.Println("SET key value - Set a key-value pair")
	fmt.Println("GET key - Retrieve the value for a given key")
	fmt.Println("DELETE key - Delete a key-value pair")
	fmt.Println("LIST - List all keys")
	fmt.Println("COUNT - Count the number of key-value pairs")
	fmt.Println("EXIT - Terminate the program")
	service.Service()
}
