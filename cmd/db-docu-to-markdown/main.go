/*
db-docu-to-markdown creates markdown files with basic information about a database.

USAGE:

	db-docu-to-markdown [global options] command [command options] [arguments...]

COMMANDS:

	help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:

	--host value, -H value                                   database host (default: "localhost")
	--port value, -P value                                   database port (default: 8080)
	--user value, -u value                                   database user (default: "admin")
	--password value, -p value                               database password (default: "password")
	--name value, -n value                                   database name (default: "test")
	--schemas value, -s value [ --schemas value, -s value ]  comma separated list of schemas to describe (default: "public")
	--dbtype value, --dt value                               specify the database type (default: "postgres")
*/
package main

import (
	"log"
	"os"

	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-docu-to-markdown/internal/generation"
	"github.com/urfave/cli/v2"
)

func main() {
	var host string
	var port int
	var user string
	var password string
	var name string
	var schemas cli.StringSlice
	var dbtype string

	app := &cli.App{
		Name:  "db-docu-to-markdown",
		Usage: "creates markdown files with basic information about a database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"H"},
				Value:       "localhost",
				Usage:       "database host",
				Destination: &host,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"P"},
				Value:       8080,
				Usage:       "database port",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:        "user",
				Aliases:     []string{"u"},
				Value:       "admin",
				Usage:       "database user",
				Destination: &user,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Value:       "password",
				Usage:       "database password",
				Destination: &password,
			},
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "test",
				Usage:       "database name",
				Destination: &name,
			},
			&cli.StringSliceFlag{
				Name:        "schemas",
				Aliases:     []string{"s"},
				Value:       cli.NewStringSlice("public"),
				Usage:       "comma separated list of schemas to describe",
				Destination: &schemas,
			},
			&cli.StringFlag{
				Name:        "dbtype",
				Aliases:     []string{"dt"},
				Value:       "postgres",
				Usage:       "specify the database type",
				Destination: &dbtype,
			},
		},
		Action: func(cCtx *cli.Context) error {
			dbDescriptorInput := connector.Input{
				Host:     host,
				Port:     port,
				User:     user,
				Password: password,
				Name:     name,
				Schemas:  schemas.Value(),
				Db:       dbtype,
			}

			return RunMarkDownGeneration(dbDescriptorInput)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func RunMarkDownGeneration(
	dbDescriptorInput connector.Input) error {
	generation.GenerateMarkdown(dbDescriptorInput)
	return nil
}
