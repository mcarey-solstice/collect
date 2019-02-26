package schema_test

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mcarey-solstice/collect/schema"
)

var yaml_str = `
---
foo:
  url: https://foo.io
  hash:
    type: sha256
    value: value
`

func testCollection(c map[string]Item, e error) {
	Expect(e).To(BeNil())

	item, ok := c["foo"]

	Expect(ok).To(BeTrue())
	Expect(item).NotTo(BeNil())
	Expect(item.Url).To(Equal("https://foo.io"))
	Expect(item.Hash.Type).To(Equal("sha256"))
	Expect(item.Hash.Value).To(Equal("value"))
}

var _ = Describe("Collection Suite", func() {
	Describe("NewCollection from string", func() {
		var file string

		BeforeEach(func() {
			tmpFile, err := ioutil.TempFile(os.TempDir(), "collection-")
			if err != nil {
				log.Fatal("Cannot create temporary file", err)
			}

			fmt.Println("Created File: " + tmpFile.Name())

			text := []byte(yaml_str)
			if _, err = tmpFile.Write(text); err != nil {
				log.Fatal("Failed to write to temporary file", err)
			}

			file = tmpFile.Name()

			// Close the file
			if err := tmpFile.Close(); err != nil {
				log.Fatal(err)
			}
		})

		AfterEach(func() {
			os.Remove(file)
		})

		It("Should load a basic configuration", func() {
			var (
				c map[string]Item
				e error
			)

			By("Using a String")
			c, e = NewCollection(yaml_str)
			testCollection(c, e)

			By("Using bytes")
			c, e = NewCollectionFromBytes([]byte(yaml_str))
			testCollection(c, e)

			By("Using a File")
			c, e = NewCollectionFromFile(file)
			testCollection(c, e)
		})
	})
})
