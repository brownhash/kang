package main

import (
	"fmt"
	"os"

	"github.com/brownhash/golog"
	"github.com/brownhash/kang/internal/command"
	"github.com/brownhash/kang/internal/core"
	"github.com/mitchellh/cli"
)

const(
	appName 	= "Kang"
	appVersion 	= "1.0.0-beta"
)

func main() {
	golog.UnsetLogFormat()

	golog.Debug(fmt.Sprintf("Initiating %v", appName))
	err := core.MaintainStructure()

	if err != nil {
		golog.Error(fmt.Errorf("Unable to initiate core structure. Error: %v", err))
	}

	golog.Debug(fmt.Sprintf("%s successfully initated", appName))

	c := cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]
	c.Commands = command.CommandCatalog()

	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err.Error())
	}

	os.Exit(exitStatus)
}