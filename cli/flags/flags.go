package flags

import "github.com/urfave/cli/v2"

func GetAvailableFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "environment",
			Aliases: []string{"e"},
			Value:   "development",
			Usage:   `Use "inFRA -e {CMD_NAME} to set the application environment.`,
		},
	}
	return flags
}
