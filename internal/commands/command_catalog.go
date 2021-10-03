package commands

import (
	"strings"

	"github.com/brownhash/kang/internal/commands/fetch"
	"github.com/brownhash/kang/internal/commands/run"
	"github.com/mitchellh/cli"
)

func CommandCatalog(commandName string, args []string) map[string]cli.CommandFactory {
	catalog := map[string]cli.CommandFactory{
		"fetch": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = fetch.Help()
			command.SynopsisText = fetch.Synopsis()

			if !strings.Contains(commandName, "help") {
				command.RunResult = fetch.Run(args)
			}
			
			return &command, nil
		},
		"run": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = run.Help()
			command.SynopsisText = run.Synopsis()

			if !strings.Contains(commandName, "help") {
				command.RunResult = run.Run(args)
			}
			
			return &command, nil
		},
	}

	if val, ok := catalog[commandName]; ok {
		return map[string]cli.CommandFactory{
			commandName: val,
		}
	}

	return catalog
}