package terraform

import (
	"os"
	"fmt"
	"path"

	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/download"
	"github.com/brownhash/session_terraform/internal/unzip"
)

func Get(version, osType, osArchitecture, downloadPath string) (string, error) {
	fileName := fmt.Sprintf("terraform_%s_%s_%s.zip", version, osType, osArchitecture)

	filePath := path.Join(downloadPath, fileName)
	cliPath := path.Join(downloadPath, "terraform")

	_, err := os.Stat(filePath)
	if err == nil {
		golog.Debug(fmt.Sprintf("terraform@%s zip already exists", version))

		_, err = os.Stat(cliPath)
		if err == nil {
			golog.Debug(fmt.Sprintf("terraform@%s cli already exists", version))
			return cliPath, err
		}

		_, err = unzip.Unzip(filePath, downloadPath)

		return cliPath, err
	}

	downloadUrl := path.Join(
		"releases.hashicorp.com/terraform", // TODO: Move to globals
		version,
		fileName,
	)

	err = download.DownloadFile("https://"+downloadUrl, downloadPath)

	if err != nil {
		return cliPath, err
	}

	_, err = unzip.Unzip(filePath, downloadPath)

	return cliPath, err
}