package cmd

import (
	"clickledger/internal/repository"
	"context"

	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

func CreateOrmCommand(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name:  "orm",
		Usage: "manage the clickledger orm",
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "start the api server",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:      "entity",
						UsageText: "name of entity to migrate",
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					var repo repository.Repository
					switch command.StringArg("entity") {
					case "link":
						repo = repository.CreateLinkRepo(db)
					case "click":
						repo = repository.CreateClickRepo(db)
					default:
						return cli.Exit("Unknown entity: "+command.StringArg("entity"), 1)
					}

					err := repo.Migrate()

					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
}
