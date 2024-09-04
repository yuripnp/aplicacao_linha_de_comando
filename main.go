package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando o programa")
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
