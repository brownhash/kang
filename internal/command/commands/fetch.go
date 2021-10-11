package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/brownhash/golog"
	"github.com/brownhash/kang/config"
	"github.com/brownhash/kang/internal/terraform"
)

// fetch command
type Fetch struct {}

func (f *Fetch) Run(args []string) int {
	if len(args) > 1 {
		golog.Error(fmt.Sprintf("Unrecognised arguments %v. fetch requires exactly 1 argument", args[1:]))
		return 1
	}

	// extract terraform version from arguments
	// latest as an argument is not supported yet
	terraformVersion := args[0]
	downloadPath := path.Join(config.SetupPath, terraformVersion)

	// generate version specific dir
	err := os.Mkdir(downloadPath, 0755)

	if err != nil {
		golog.Debug(err)
	}

	// get cli path of downloaded terraform cli
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
	return `
Fetch a specific version of terraform by giving version.
kang fetch <version>
	version: 0.14.5 OR 1.0.0 OR ...
`
}

func (f *Fetch) Synopsis() string {
	return "Fetch a specific terraform version."
}
