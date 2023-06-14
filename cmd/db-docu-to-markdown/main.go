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
		Name:  "db-docu-sync-confluence",
		Usage: "syncronises one or more pages in confluence with database documentation",
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
