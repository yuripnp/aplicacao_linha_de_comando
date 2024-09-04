package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	// exemplo comand: go run main.go ip --host exemplo.com.br
	fmt.Println("Iniciando o programa")
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
