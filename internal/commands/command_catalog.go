package commands

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/brownhash/session_terraform/internal/core"
	"github.com/brownhash/session_terraform/internal/commands/session_details"
	"github.com/brownhash/session_terraform/internal/commands/fetch"
	"github.com/brownhash/session_terraform/internal/commands/run"
)

func CommandCatalog(s core.Session, commandName string, args []string) map[string]cli.CommandFactory {
	catalog := map[string]cli.CommandFactory{
		"session": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = session_details.Help()
			command.SynopsisText = session_details.Synopsis()

			if !strings.Contains(commandName, "help") {
				command.RunResult = session_details.Run(s)
			}
			
			return &command, nil
		},
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