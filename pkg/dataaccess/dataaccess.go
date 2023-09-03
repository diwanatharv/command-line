package dataaccess

import (
	"awesomeProject1/pkg/models"
	"fmt"
	"os"
	"strings"
)

var store map[string]string
var dataFilePath = "data.txt"

func saveData() {
	store = make(map[string]string)
	file, err := os.Create(dataFilePath)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	defer file.Close()

	for key, value := range store {
		line := fmt.Sprintf("%s=%s\n", key, value)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing data:", err)
			return
		}
	}

	return
}

func ProcessCommand(input string, kvs *models.KeyValueStore) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := strings.ToUpper(parts[0])
	switch command {
	case "SET":
		if len(parts) != 3 {
			fmt.Println("Invalid SET command format. Use SET key value.")
			return
		}
		key := parts[1]
		value := parts[2]

		kvs.Set(key, value) // Call the Set method on the KeyValueStore
		saveData()
		fmt.Printf("Key \"%s\" set successfully.\n", key)

	case "GET":
		if len(parts) != 2 {
			fmt.Println("Invalid GET command format. Use GET key.")
			return
		}
		key := parts[1]

		value, exists := kvs.Get(key) // Call the Get method on the KeyValueStore
		if exists {
			fmt.Printf("Value for key \"%s\": %s\n", key, value)
		} else {
			fmt.Printf("Key \"%s\" not found.\n", key)
		}

	case "DELETE":
		if len(parts) != 2 {
			fmt.Println("Invalid DELETE command format. Use DELETE key.")
			return
		}
		key := parts[1]

		deleted := kvs.Delete(key) // Call the Delete method on the KeyValueStore
		if deleted {
			saveData()
			fmt.Printf("Key \"%s\" deleted successfully.\n", key)
		} else {
			fmt.Printf("Key \"%s\" not found.\n", key)
		}

	case "LIST":
		keys := kvs.ListKeys() // Call the ListKeys method on the KeyValueStore
		fmt.Println("Keys in the store:")
		for _, key := range keys {
			fmt.Println(key)
		}

	case "COUNT":
		count := kvs.Count() // Call the Count method on the KeyValueStore
		fmt.Printf("Number of key-value pairs in the store: %d\n", count)

	case "EXIT":
		saveData()
		fmt.Println("Goodbye!")
		os.Exit(0)

	default:
		fmt.Println("Unknown command:", command)
	}
}
