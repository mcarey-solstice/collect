package file_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mcarey-solstice/collect/file"
)

var _ = Describe("File", func() {

	Describe("File existence", func() {
		var (
			tmp string
			err error
		)

		BeforeEach(func() {
			tmp, err = ioutil.TempDir(os.TempDir(), "collect-")
			Expect(err).To(BeNil())
		})

		It("Should know the file exists", func() {
			file, err := os.Create(filepath.Join(tmp, "there"))
			Expect(err).To(BeNil())

			ok := Exists(file.Name())
			Expect(ok).To(BeTrue())
		})

		It("Should know the file does not exist", func() {
			ok := Exists(filepath.Join(tmp, "nothere"))
			Expect(ok).To(BeFalse())
		})
	})
})
