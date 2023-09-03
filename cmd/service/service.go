package service

import (
	"awesomeProject1/pkg/dataaccess"
	"awesomeProject1/pkg/models"
	"bufio"
	"fmt"
	"os"
)

func Service() {
	kvs := models.NewKeyValueStore()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		dataaccess.ProcessCommand(input, kvs)
	}
}
