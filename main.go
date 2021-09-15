package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"sample": func() (cli.Command, error) {

			var command cli.MockCommand
			command.HelpText = HelpCommand()
			command.RunResult = RunCommand([]string{"random argument"})
			command.SynopsisText = SynopsisCommand()
			
			return &command, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
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