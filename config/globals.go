package config

import (
	"os"
	"path"
)

var TerraformRegistryUrl string = "releases.hashicorp.com/terraform"
var SetupPath string = path.Join(os.Getenv("HOME"), ".sessionTerraform")