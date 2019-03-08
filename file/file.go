package file

import (
	"fmt"
	"os"

	"github.com/mcarey-solstice/collect/schema"
)

func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

func CollectAll(collection map[string]schema.Item) []error {
	errs := []error{}
	for key, value := range collection {
		err := Collect(key, &value)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func Collect(file string, item *schema.Item) error {
	if Exists(file) {
		// If the file exists, hash it
		h, e := HashFile(item.Hash.Type, file)
		if e != nil {
			return e
		}
		if h != item.Hash.Value {
			// If the hash does not match, download it
			fmt.Printf("Hash of %s does not match. Redownloading from: %s\n", file, item.Url)
			fmt.Printf("%s != %s", h, item.Hash.Value)
			return downloadFile(item.Url, file, item.Mode)
		}
	} else {
		fmt.Printf("%s does not exist. Downloading from: %s\n", file, item.Url)
		return downloadFile(item.Url, file, item.Mode)
	}

	return nil
}

func downloadFile(url string, file string, mode os.FileMode) error {
	var err error

	DownloadFile(url, file)
	if mode != 0 {
		err = os.Chmod(file, mode)
	}

	return err
}
