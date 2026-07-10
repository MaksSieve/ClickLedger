package main

import (
	"context"
	"log"
	"os"

	"clickledger/cmd"
)

func main() {
	command, err := cmd.CreateCli()

	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
