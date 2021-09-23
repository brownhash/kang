package config

import (
	"os"
	"path"
)

var TerraformRegistryUrl string = "releases.hashicorp.com/terraform"
var SetupPath string = path.Join(os.Getenv("HOME"), ".sessionTerraform")
var DefaultDirPermission os.FileMode = 755
var TempExtension string = ".tmp"
var OsType string = "darwin" // TODO: Find for system
var OsArchitecture string = "amd64" // TODO: Find for system