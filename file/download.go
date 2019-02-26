package file

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func DownloadFile(uri string, filepath string) error {
	u, e := url.Parse(uri)
	if e != nil {
		return e
	}

	switch u.Scheme {
	case "file":
		return FileGet(u, filepath)
	case "http":
		return HttpGet(u, filepath)
	case "https":
		return HttpGet(u, filepath)
	default:
		return errors.New(fmt.Sprintf("Unknown scheme: %s", u.Scheme))
	}
}

func FileGet(uri *url.URL, filepath string) error {
	fmt.Printf("Fetching data from file: %s\n", uri)

	fmt.Printf("Creating file: %s\n", filepath)
	// Create the file
	out, err := os.Create(filepath)
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

func HttpGet(uri *url.URL, filepath string) error {
	fmt.Printf("Fetching data from http: %s\n", uri)
	// Get the data
	resp, err := http.Get(uri.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Creating file: %s\n", filepath)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Copying file out")
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
