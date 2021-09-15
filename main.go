package main

import (
	"fmt"
	"os"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/session"
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
		"sample": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = HelpCommand()
			command.RunResult = RunCommand(c.Args)
			command.SynopsisText = SynopsisCommand()
			
			return &command, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err.Error())
	}

	os.Exit(exitStatus)
}

func RunCommand(args []string) int {
	return 0
}

func HelpCommand() string {
	return "Help text"
}

func SynopsisCommand() string {
	return "Synopsis text"
}