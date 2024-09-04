package app

import "github.com/urfave/cli"

// vai ser em maiusculo pq eu quero que seja exportado para a main
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e nomes de servidores"

	return app
}
