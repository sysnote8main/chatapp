package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"chatapp/client"
	"chatapp/server"
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
					port := ctx.Int("port")
					server.Run(port)
					return nil
				},
			},
			{
				Name: "client",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   3000,
					},
				},
				Action: func(ctx *cli.Context) error {
					port := ctx.Int("port")
					client.Run(port)
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
