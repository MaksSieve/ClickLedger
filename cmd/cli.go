package cmd

import (
	"log"

	"github.com/urfave/cli/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateCli() (*cli.Command, error) {

	dsn := "host=localhost user=clickledger password=clickledger dbname=clickledger port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	log.Default().Println("Database connected")

	return &cli.Command{
		Name:        "clickledger",
		Aliases:     []string{"cl"},
		Description: "ClickLedger is a simple and efficient ledger application that allows users to manage their financial transactions with ease. It provides a command-line interface for adding, viewing, and managing transactions, as well as generating reports and summaries.",
		Commands: []*cli.Command{
			CreateApiCommand(db),
			CreateOrmCommand(db),
		},
	}, nil
}
