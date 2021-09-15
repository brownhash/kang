package main

import (
	"fmt"
	"os"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/session"
	"github.com/brownhash/session_terraform/internal/commands/session_details"
	"github.com/mitchellh/cli"
)

const(
	appName 	= "SessionTerraform"
	appVersion 	= "1.0.0"
)

func main() {
	golog.Info(fmt.Sprintf("Initiating %v", appName))
	session := session.GenerateSession()

	golog.Success(fmt.Sprintf("%s initated for %s at %v", appName, session.User, session.Started))

	c := cli.NewCLI(appName, appVersion)

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"session": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = session_details.Help()
			command.RunResult = session_details.Run(session)
			command.SynopsisText = session_details.Synopsis()
			
			return &command, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err.Error())
	}

	os.Exit(exitStatus)
}
