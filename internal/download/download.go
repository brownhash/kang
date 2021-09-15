package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/brownhash/golog"
)

func File(fileUrl, downloadPath string) {
	// Build fileName from fullPath
	fileURL, err := url.Parse(fileUrl)
	if err != nil {
		log.Fatal(err)
	}

	urlPath := fileURL.Path
	segments := strings.Split(urlPath, "/")
	fileName := segments[len(segments)-1]

	golog.Info(fmt.Sprintf("Downloading %s from %s", fileName, fileURL))

	// Create blank file
	file, err := os.Create(path.Join(downloadPath, fileName))
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	defer file.Close()

	golog.Success(fmt.Sprintf("%s Downloaded", fileName))
}