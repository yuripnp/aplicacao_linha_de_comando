package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// vai ser em maiusculo pq eu quero que seja exportado para a main
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e nomes de servidores"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
