package main

import (
	"cli/commands"
	"cli/commands/flags"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:     "infra",
		Flags:    flags.GetGlobalFlags(),
		Commands: commands.GetAllCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
