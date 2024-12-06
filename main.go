package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "chatapp",
		Usage: "Let's chat!",
		Commands: []*cli.Command{
			{
				Name: "server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   3000,
					},
				},
				Action: func(ctx *cli.Context) error {
					fmt.Println("Server Hogehoge")
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
