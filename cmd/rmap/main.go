package main

import (
	"fmt"
	"log"
	"os"

	"github.com/reconmap/cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.App{
		Name: "Reconmap CLI",
		Authors: []*cli.Author{
			{
				Name:  "Reconmap developers",
				Email: "devs@reconmap.org",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "login",
				Usage: "initiates a session with the server",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "username", Aliases: []string{"u"}, Required: true},
					&cli.StringFlag{Name: "password", Aliases: []string{"p"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					response, err := commands.Login(c.String("username"), c.String("password"))
					if err == nil {
						fmt.Printf("%s\n", response)
					}
					return err
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "runs a command and upload the results",
				Action: func(c *cli.Context) error {
					_, err := commands.CreateNewContainer("reconmap/pentest-container-tools-goohost")
					//_, err := commands.CreateNewContainer("nginx")
					return err
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list Reconmap containers",
				Action: func(c *cli.Context) error {
					commands.ListContainer()
					return nil
				},
			},
			{
				Name:    "upload-results",
				Aliases: []string{"u"},
				Usage:   "upload command results",
				Action: func(c *cli.Context) error {
					err := commands.UploadResults()
					if err != nil {
						fmt.Printf("%s\n", err)
					}
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
