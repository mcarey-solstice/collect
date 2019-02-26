package file

import (
	"fmt"
	"os"

	"github.com/mcarey-solstice/collect/schema"
)

const (
	DEFAULT_COLLECTION_DIRECTORY = ".collect"
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
			DownloadFile(item.Url, file)
		}
	} else {
		fmt.Printf("%s does not exist. Downloading from: %s\n", file, item.Url)
		DownloadFile(item.Url, file)
	}

	return nil
}
