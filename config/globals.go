package config

import (
	"os"
	"path"
	"runtime"
)

var TerraformRegistryUrl string = "releases.hashicorp.com/terraform"
var SetupPath string = path.Join(os.Getenv("HOME"), ".sessionTerraform")
var DefaultDirPermission os.FileMode = 755
var TempExtension string = ".tmp"
var OsType string = runtime.GOOS
var OsArchitecture string = runtime.GOARCH
