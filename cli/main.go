package main

import (
	"cli/flags"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Flags: flags.GetAvailableFlags(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
