package fetch

import (
	"os"
	"fmt"
	"path"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/terraform"
)

func Run(args []string) int {
	setupPath := os.Getenv("HOME") // TODO: Move to globals
	setupPath = path.Join(setupPath, ".sessionTerraform") // TODO: Move to globals

	if len(args) > 1 {
		golog.Error(fmt.Sprintf("Unrecognised arguments %v. fetch requires exactly 1 argument", args[1:]))
		return 1
	}

	terraformVersion := args[0]

	osType := "darwin" // TODO: Move to globals
	osArch := "amd64" // TODO: Move to globals
	downloadPath := path.Join(setupPath, terraformVersion)

	err := os.Mkdir(downloadPath, 0755)

	if err != nil {
		golog.Debug(err)
	}

	cliPath, err := terraform.Get(terraformVersion, osType, osArch, downloadPath)

	if err != nil {
		golog.Error(fmt.Sprintf("Unable to download terraform@%s. Error: %v", terraformVersion, err))
		return 1
	}
	golog.Success(fmt.Sprintf("terraform@%s successfully fetched!", terraformVersion))
	golog.Debug(fmt.Sprintf("terraform@%s stored at %s", terraformVersion, cliPath))

	return 0
}

func Help() string {
	return "\n\tFetch a specific version of terraform by giving 0.14.7.\n"
}

func Synopsis() string {
	return "Fetch a terraform version."
}