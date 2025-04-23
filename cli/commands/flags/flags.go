package flags

import "github.com/urfave/cli/v2"

func GetCreateActionFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "registry",
			Aliases: []string{"r"},
			Value:   "access",
			Usage:   `Use "inFRA {SUBCOMMAND} -r {REGISTRY_FILE_NAME} to create a windows registry for the db connection.`,
		},
		&cli.StringFlag{
			Name:    "container-name",
			Aliases: []string{"cn"},
			Value:   "container",
			Usage:   `Use "inFRA {SUBCOMMAND} -name {CONTAINER_NAME} to set the name for the Docker Container.`,
		},
	}
}

func GetGlobalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "environment",
			Aliases: []string{"e"},
			Value:   "development",
			Usage:   `Use "inFRA {SUBCOMMAND} -e {ENV_NAME} to set the application environment.`,
		},
	}
}
