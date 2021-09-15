package setup

import (
	"fmt"
	"os"
	"path"

	"github.com/brownhash/golog"
)

func Run(args []string) int {
	setupPath := os.Getenv("HOME")

	if len(args) > 0 {
		golog.Error(fmt.Sprintf("Unrecognised arguments %v. setup requires no argument", args))
		return 1
	}

	setupPath = path.Join(setupPath, ".sessionTerraform")

	_, err := os.Stat(setupPath)
	if err == nil {
		golog.Warn(fmt.Sprintf("Set up already done at %v", setupPath))
		return 0
	}

	err = os.Mkdir(setupPath, 0755)

	if err != nil {
		golog.Error(err)
		return 1
	}

	golog.Success(fmt.Sprintf("Set up complete at %v", setupPath))

	return 0
}

func Help() string {
	return "\n\tSets up required structure and dependencies for self.\n"
}

func Synopsis() string {
	return "Self setup."
}
