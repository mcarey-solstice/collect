package schema

import (
	"log"
	"fmt"
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Hash struct {
	Type string `yaml:"type"`
	Value string `yaml:"value"`
}

type Item struct {
	Url string `yaml:"url"`
	Hash Hash `yaml:"hash"`
	Unpack bool `yaml:"unpack"`
}

func VerifyCollection(collection map[string]Item) error {
	for k, v := range collection {
		if v.Hash.Type != "" && v.Hash.Value == "" {
			return errors.New(fmt.Sprintf("Cannot specify a hash type without a value: item %s", k))
		}
	}

	return nil
}

func NewCollectionFromBytes(bytes []byte) (map[string]Item, error) {
	c := make(map[string]Item)

	err := yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
		return c, err
	}

	if e := VerifyCollection(c); e != nil {
		log.Fatalf("error: %v", e)
		return c, e
	}

	return c, nil
}

func NewCollection(s string) (map[string]Item, error) {
	return NewCollectionFromBytes([]byte(s))
}

func NewCollectionFromFile(filename string) (map[string]Item, error) {
	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
		c := make(map[string]Item)
		return c, err
	}

	return NewCollectionFromBytes(yml)
}
