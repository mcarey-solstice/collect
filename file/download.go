package file

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func DownloadFile(uri string, path string) error {
	u, e := url.Parse(uri)
	if e != nil {
		return e
	}

	// Create the directory needed
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	switch u.Scheme {
	case "file":
		return FileGet(u, path)
	case "http":
		return HttpGet(u, path)
	case "https":
		return HttpGet(u, path)
	default:
		return errors.New(fmt.Sprintf("Unknown scheme: %s", u.Scheme))
	}
}

func FileGet(uri *url.URL, path string) error {
	fmt.Printf("Fetching data from file: %s\n", uri)

	fmt.Printf("Creating file: %s\n", path)
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(uri.Path)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	return err
}

func HttpGet(uri *url.URL, path string) error {
	fmt.Printf("Fetching data from http: %s\n", uri)
	// Get the data
	resp, err := http.Get(uri.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Creating file: %s\n", path)
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Copying file out")
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
