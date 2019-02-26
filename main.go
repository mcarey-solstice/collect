package main

import (
	"log"
	"os"

	. "github.com/mcarey-solstice/collect/file"
	. "github.com/mcarey-solstice/collect/schema"
)

const DEFAULT_FILE = "collect.yml"

func main() {
	file := DEFAULT_FILE

	args := os.Args[1:]
	if len(args) > 0 {
		file = args[0]
	}

	c, e := NewCollectionFromFile(file)
	if e != nil {
		log.Fatalf("%v", e)
		panic("Failed to create collection from file")
	}

	if err := CollectAll(c); err != nil {
		log.Fatalf("%v", err)
		panic("Failed to collect all")
	}
}
