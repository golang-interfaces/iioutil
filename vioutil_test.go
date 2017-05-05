package vioutil

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/virtual-go/fs"
	"github.com/virtual-go/fs/osfs"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

var _ = Context("_VIOUtil", func() {
	wd, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	Context("New", func() {
		It("should return Fs", func() {
			/* arrange/act/assert */
			Expect(New(new(fs.Fake))).
				Should(Not(BeNil()))
		})
	})
	Context("ReadDir", func() {
		It("should return expected fileinfos", func() {
			/* arrange */
			// use .opspec dir because it won't be modified during test
			providedDirName := path.Join(wd, ".opspec")

			expectedFileInfos, _ := ioutil.ReadDir(providedDirName)

			objectUnderTest := _VIOUtil{
				fs: osfs.New(),
			}

			/* act */
			actualFileinfos, actualErr := objectUnderTest.ReadDir(providedDirName)

			/* assert */
			Expect(actualFileinfos).To(Equal(expectedFileInfos))
			Expect(actualErr).To(BeNil())
		})
	})
	Context("ReadFile", func() {
		It("should return expected fileinfo", func() {
			/* arrange */
			// use current file for test
			_, providedFileName, _, _ := runtime.Caller(1)

			expectedBytes, _ := ioutil.ReadFile(providedFileName)

			objectUnderTest := _VIOUtil{
				fs: osfs.New(),
			}

			/* act */
			actualBytes, actualErr := objectUnderTest.ReadFile(providedFileName)

			/* assert */
			Expect(actualBytes).To(Equal(expectedBytes))
			Expect(actualErr).To(BeNil())
		})
	})
	Context("WriteFile", func() {
		It("should create expected file", func() {
			/* arrange */
			tempFile, err := ioutil.TempFile("", "dummyFile")
			if nil != err {
				panic(err)
			}
			providedFilename := tempFile.Name()

			providedData := bytes.NewBufferString("dummy file content").Bytes()
			providedPerm := os.FileMode(0777)

			objectUnderTest := _VIOUtil{
				fs: osfs.New(),
			}

			/* act */
			objectUnderTest.WriteFile(providedFilename, providedData, providedPerm)

			/* assert */
			actualData, err := ioutil.ReadFile(providedFilename)
			if nil != err {
				panic(err)
			}

			Expect(actualData).To(Equal(providedData))
		})
	})
})