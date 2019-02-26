package file

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	// SHA1   = "sha1"
	SHA256 = "sha256"
	// MD5    = "md5"
)

var HASH_TYPES = [1]string{SHA256} // [3]string{SHA1, SHA256, MD5}

func HashFile(t string, file string) (string, error) {
	f, e := os.Open(file)
	fmt.Println(f.Name())
	if e != nil {
		return "", e
	}
	defer f.Close()

	switch t {
	case SHA256:
		return HashFileWithSha256(f)
	default:
		return "", errors.New(fmt.Sprintf("Unknown hash type: %s, expected one of: %v", t, HASH_TYPES))
	}
}

func HashFileWithSha256(file *os.File) (string, error) {
	hash := sha256.New()
	if _, e := io.Copy(hash, file); e != nil {
		return "", e
	}

	return string(hash.Sum(nil)), nil
}
