package main

import (
	"fmt"
	"os"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/session"
	"github.com/brownhash/session_terraform/internal/commands"
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

	commandName := os.Args[1]
	args := []string{}

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	c := cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]
	c.Commands = commands.CommandCatalog(session, commandName, args)

	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err.Error())
	}

	os.Exit(exitStatus)
}
