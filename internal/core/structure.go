package core

import (
	"os"

	"github.com/brownhash/kang/config"
)

func MaintainStructure() error {
	_, err := os.Stat(config.SetupPath)

	if err == nil {
		return err
	}

	err = os.Mkdir(config.SetupPath, config.DefaultDirPermission)

	return err
}
