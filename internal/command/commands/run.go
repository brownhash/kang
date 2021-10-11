package commands

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/brownhash/golog"
	"github.com/brownhash/kang/config"
	"github.com/brownhash/kang/internal/shell"
	"github.com/brownhash/kang/internal/terraform"
)

// Run command
type RunCommand struct {}

func (r *RunCommand) Run(args []string) int {
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

	// extract terraform command to run from arguments
	terraformCommand := args[1]

	terraformArgs := []string{}
	if len(args) > 1 {
		terraformArgs = args[2:]
	}

	return shell.Exec(fmt.Sprintf("%s %s %s", cliPath, terraformCommand, strings.Join(terraformArgs, " ")))
}

func (r *RunCommand) Help() string {
	return `
Run a specific version of terraform by giving terraform version and args.
kang run <version> <args>
	version: 0.14.5/1.0.0/...etc
	args: plan -var-file=terraform.tfvars -out=plan.out OR apply plan.out OR ...
`
}

func (r *RunCommand) Synopsis() string {
	return "Run a terraform command with a specific version."
}