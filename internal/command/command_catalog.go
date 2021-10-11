package command

import (
	"github.com/brownhash/kang/internal/command/commands"
	"github.com/mitchellh/cli"
)

func CommandCatalog() map[string]cli.CommandFactory {
	catalog := map[string]cli.CommandFactory{
		"fetch": func() (cli.Command, error) {
			return &commands.Fetch{}, nil
		},
		"run": func() (cli.Command, error) {
			return &commands.RunCommand{}, nil
		},
	}

	return catalog
}