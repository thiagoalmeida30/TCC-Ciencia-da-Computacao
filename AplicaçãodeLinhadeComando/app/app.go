package app

import "github.com/urfave/cli"

// gerar retorna a função pronta para executar a linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação delinha de comando"
	app.Usage = "busca Ips e Nomes de Servidor na internet"

	return app
}
