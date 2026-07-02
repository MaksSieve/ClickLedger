package main

import (
	"context"
	"log"
	"os"

	"clickledger/internal/router"

	"github.com/urfave/cli/v3"
)

func main() {
	command := &cli.Command{
		Name: "clickledger",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Value:   "8080",
				Usage:   "port for the api",
				Aliases: []string{"p"},
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			port := command.String("port")
			router.API(port)
			return nil
		},
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
