package commands

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/brownhash/session_terraform/internal/session"
	"github.com/brownhash/session_terraform/internal/commands/setup"
	"github.com/brownhash/session_terraform/internal/commands/session_details"
)

func CommandCatalog(s session.Session, commandName string, args []string) map[string]cli.CommandFactory {
	catalog := map[string]cli.CommandFactory{
		"setup": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = setup.Help()
			command.SynopsisText = setup.Synopsis()

			if !strings.Contains(commandName, "help") {
				command.RunResult = setup.Run(args)
			}
			
			return &command, nil
		},
		"session": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = session_details.Help()
			command.SynopsisText = session_details.Synopsis()

			if !strings.Contains(commandName, "help") {
				command.RunResult = session_details.Run(s)
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