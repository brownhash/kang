package main

import (
	"fmt"
	"os"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/core"
	"github.com/brownhash/session_terraform/internal/commands"
	"github.com/mitchellh/cli"
)

const(
	appName 	= "Kang"
	appVersion 	= "1.0.0"
)

func main() {
	golog.Info(fmt.Sprintf("Initiating %v", appName))
	session := core.GenerateSession()
	err := core.MaintainStructure()

	if err != nil {
		golog.Error(fmt.Errorf("Unable to initiate core structure. Error: %v", err))
	}

	golog.Success(fmt.Sprintf("%s initated for %s at %v", appName, session.User, session.Started))

	commandName := "--help"
	if len(os.Args) > 1 {
		commandName = os.Args[1]
	}
	
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	c := cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]

	session.Command = commandName
	session.Arguments = args
	c.Commands = commands.CommandCatalog(session, commandName, args)

	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err.Error())
	}

	defer session.Save(exitStatus)

	os.Exit(exitStatus)
}