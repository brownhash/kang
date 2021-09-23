package setup

import (
	"fmt"
	"os"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/config"
)

func Run(args []string) int {
	if len(args) > 0 {
		golog.Error(fmt.Sprintf("Unrecognised arguments %v. setup requires no argument", args))
		return 1
	}

	_, err := os.Stat(config.SetupPath)
	if err == nil {
		golog.Debug(fmt.Sprintf("Set up already done at %v", config.SetupPath))
		return 0
	}

	err = os.Mkdir(config.SetupPath, 0755)

	if err != nil {
		golog.Error(err)
		return 1
	}

	golog.Success(fmt.Sprintf("Set up complete at %v", config.SetupPath))

	return 0
}

func Help() string {
	return "\n\tSets up required structure and dependencies for self.\n"
}

func Synopsis() string {
	return "Self setup."
}
