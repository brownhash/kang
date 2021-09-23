package terraform

import (
	"os"
	"fmt"
	"path"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/download"
)

func Get(version, osType, osArchitecture, downloadPath string) (string, error) {
	fileName := fmt.Sprintf("terraform_%s_%s_%s.zip", version, osType, osArchitecture)

	filePath := path.Join(downloadPath, fileName)
	cliPath := path.Join(downloadPath, "terraform")

	_, err := os.Stat(filePath)
	if err == nil {
		golog.Warn(fmt.Sprintf("terraform@%s zip already exists", version))

		_, err = os.Stat(cliPath)
		if err == nil {
			golog.Warn(fmt.Sprintf("terraform@%s cli already exists", version))
			return cliPath, err
		}

		// TODO: extract filePath at downloadPath

		return cliPath, err
	}

	downloadUrl := path.Join(
		"https://releases.hashicorp.com/terraform",
		version,
		fileName,
	)

	err = download.DownloadFile(downloadUrl, downloadPath)

	// TODO: extract filePath at downloadPath

	return cliPath, err
}

