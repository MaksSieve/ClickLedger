package cmd

import (
	"clickledger/internal/router"
	"context"
	"log"
	"net/http"

	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

func CreateApiCommand(db *gorm.DB) *cli.Command {

	return &cli.Command{
		Name:  "api",
		Usage: "manage the clickledger api server",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start the api server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "addr",
						Value:   "localhost",
						Usage:   "address for the api",
						Aliases: []string{"a"},
					},
					&cli.StringFlag{
						Name:    "port",
						Value:   "8080",
						Usage:   "port for the api",
						Aliases: []string{"p"},
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					port := command.String("port")
					addr := command.String("addr")

					if addr == "" || port == "" {
						return nil
					}
					addrString := addr + ":" + port
					log.Println("server started at", addrString)
					router := router.CreateMainRouter(db)
					http.ListenAndServe(addrString, router)
					return nil
				},
			},
		},
	}
}
