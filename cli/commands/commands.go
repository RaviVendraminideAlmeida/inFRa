package commands

import (
	"cli/commands/actions"
	"cli/commands/flags"

	"github.com/urfave/cli/v2"
)

func GetAllCommands() []*cli.Command {
	commands := []*cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a new environment for the application.",
			Action:  actions.CreateAction,
			Flags:   flags.GetCreateActionFlags(),
		},
	}
	return commands
}
