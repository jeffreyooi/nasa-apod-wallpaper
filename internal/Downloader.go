package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

const (
	// BaseURL ...
	BaseURL = "https://api.nasa.gov"

	// APIAPOD ...
	APIAPOD = "/planetary/apod"
)

const (
	// ErrAPIKeyEmpty ...
	ErrAPIKeyEmpty = "API key empty"
	// ErrAPIKeyDemo ...
	ErrAPIKeyDemo = "using demo API key"
	// ErrAPIKeyNotSet ...
	ErrAPIKeyNotSet = "API key not set"
	// ErrAPIKeyInvalid ...
	ErrAPIKeyInvalid = "invalid API key"
)

// DownloadAPOD ...
func DownloadAPOD(apiKey string, date string, requestHD bool) (string, error) {
	if apiKey == "" {
		return "", errors.New(ErrAPIKeyEmpty)
	}
	if apiKey == "DEMO_KEY" {
		return "", errors.New(ErrAPIKeyDemo)
	}
	if apiKey == "<Paste your API Key here>" {
		return "", errors.New(ErrAPIKeyNotSet)
	}

	parsedURL, err := url.Parse(BaseURL)
	if err != nil {
		return "", err
	}
	parsedURL.Path = APIAPOD

	query := parsedURL.Query()
	query.Add("api_key", apiKey)
	if date != "" {
		query.Add("date", date)
	}
	parsedURL.RawQuery = query.Encode()

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 403 {
		return "", errors.New(ErrAPIKeyInvalid)
	}
	if resp.StatusCode != 200 {
		return "", errors.Errorf("failed to fetch APOD, status code = %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var apodResponse APODResponse
	if err := json.NewDecoder(resp.Body).Decode(&apodResponse); err != nil {
		return "", err
	}

	var downloadURL string
	if requestHD {
		downloadURL = apodResponse.HDURL
	} else {
		downloadURL = apodResponse.URL
	}

	return downloadFile(downloadURL)
}

func downloadFile(downloadURL string) (string, error) {
	imgResp, err := http.Get(downloadURL)
	if err != nil {
		return "", err
	}
	if imgResp.StatusCode != 200 {
		return "", errors.Errorf("failed to download APOD, status code = %d", imgResp.StatusCode)
	}
	defer imgResp.Body.Close()

	url, _ := url.Parse(downloadURL)
	urlPaths := strings.Split(url.Path, "/")
	fileName := urlPaths[len(urlPaths)-1]

	outputPath := path.Join(os.TempDir(), fileName)

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, imgResp.Body)
	return outputPath, err
}
