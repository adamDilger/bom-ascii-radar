package bom

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	imgUrlFormat = "http://www.bom.gov.au/%s"
)

type FetchImageOptions struct {
	TheImageName string
	CacheDir     string
}

func GetRadarImagePath(opts FetchImageOptions) (string, error) {
	name := opts.TheImageName
	cacheDir := opts.CacheDir

	slog.Debug("Fetching image", "theImageName", name)

	url := fmt.Sprintf(imgUrlFormat, name)

	p := path.Join(cacheDir, name)

	if fileExists(p) {
		slog.Debug("using cached image", "url", p)
		return p, nil
	}

	os.MkdirAll(filepath.Dir(p), os.ModePerm)
	f, err := os.Create(p)
	if err != nil {
		panic("Failed to open cache dir: " + err.Error())
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for url [%s]: %v", url, err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("failed to fetch immage for url [%s]: %v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to read image for product code [%s]: %v", name, resp.Status)
	}

	slog.Debug("saving image to cache", "url", p)

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body into cache for url [%s]: %v", url, err)
	}

	return p, nil
}

func GetDateTimeForImagePath(path string) (time.Time, error) {
	// radar/IDR763.T.202403050758.png
	// len(path) - len("202403050758") - len(".png")
	timestamp := path[len(path)-12-4 : len(path)-4]

	return time.Parse("200601021504", timestamp)
}

func DeleteImage(imageName string, cacheDir string) {
	p := path.Join(cacheDir, imageName)
	if !fileExists(p) {
		slog.Debug("No cached image to delete", "url", p)
		return
	}

	slog.Debug("deleting image", "url", p)
	if err := os.Remove(p); err != nil {
		slog.Debug("failed to delete image", "url", p)
	}
}

func fileExists(p string) bool {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		return false
	}

	return true
}
