package main

import (
	"fmt"

	"github.com/priscila-albertini/singleton/internal"
)

func main() {
	logger := internal.GetLogger()

	logger.Log("Primeira mensagem de log.")
	logger.Log("Segunda mensagem de log.")

	// Verificando se há apenas uma única instância de Logger
	anotherLogger := internal.GetLogger()

	if logger == anotherLogger {
		fmt.Println("Apenas uma única instância do Logger está sendo utilizada.")
	}
}
