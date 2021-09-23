package run

import (
	"os"
	"fmt"
	"path"
	"strings"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/terraform"
	"github.com/brownhash/session_terraform/internal/shell"
	"github.com/brownhash/session_terraform/config"
)

func Run(args []string) int {
	terraformVersion := args[0]
	downloadPath := path.Join(config.SetupPath, terraformVersion)

	err := os.Mkdir(downloadPath, 0755)

	if err != nil {
		golog.Debug(err)
	}

	cliPath, err := terraform.Get(terraformVersion, config.OsType, config.OsArchitecture, downloadPath)

	terraformCommand := args[1]

	terraformArgs := []string{}
	if len(args) > 1 {
		terraformArgs = args[2:]
	}

	return shell.Exec(fmt.Sprintf("%s %s %s", cliPath, terraformCommand, strings.Join(terraformArgs, " ")))
}

func Help() string {
	return "\n\tRun a specific version of terraform by giving terraform version and args.\n"
}

func Synopsis() string {
	return "Run a terraform version."
}