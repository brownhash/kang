package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/brownhash/golog"
	"github.com/brownhash/kang/config"
	"github.com/brownhash/kang/internal/terraform"
)

type Fetch struct {}

func (f *Fetch) Run(args []string) int {
	if len(args) > 1 {
		golog.Error(fmt.Sprintf("Unrecognised arguments %v. fetch requires exactly 1 argument", args[1:]))
		return 1
	}

	terraformVersion := args[0]
	downloadPath := path.Join(config.SetupPath, terraformVersion)

	err := os.Mkdir(downloadPath, 0755)

	if err != nil {
		golog.Debug(err)
	}

	cliPath, err := terraform.Get(terraformVersion, config.OsType, config.OsArchitecture, downloadPath)

	if err != nil {
		golog.Error(fmt.Sprintf("Unable to download terraform@%s. Error: %v", terraformVersion, err))
		return 1
	}
	golog.Success(fmt.Sprintf("terraform@%s successfully fetched!", terraformVersion))
	golog.Debug(fmt.Sprintf("terraform@%s stored at %s", terraformVersion, cliPath))

	return 0
}

func (f *Fetch) Help() string {
	return "\n\tFetch a specific version of terraform by giving 0.14.7.\n"
}

func (f *Fetch) Synopsis() string {
	return "Fetch a terraform version."
}