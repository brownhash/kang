package config

import (
	"os"
	"path"
	"runtime"
)

// global configurations
var TerraformDownloadUrl string = "https://releases.hashicorp.com/terraform/%s/%s"
var SetupPath string = path.Join(os.Getenv("HOME"), ".kang")
var DefaultDirPermission os.FileMode = 755
var TempExtension string = ".tmp"
var OsType string = runtime.GOOS
var OsArchitecture string = runtime.GOARCH
