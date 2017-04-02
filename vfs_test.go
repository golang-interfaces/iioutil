package vioutil

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
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
			Expect(New()).Should(Not(BeNil()))
		})
	})
	Context("ReadDir", func() {
		It("should return expected fileinfos", func() {
			/* arrange */
			providedDirName := wd

			expectedFileInfos, _ := ioutil.ReadDir(providedDirName)

			objectUnderTest := New()

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

			objectUnderTest := New()

			/* act */
			actualBytes, actualErr := objectUnderTest.ReadFile(providedFileName)

			/* assert */
			Expect(actualBytes).To(Equal(expectedBytes))
			Expect(actualErr).To(BeNil())
		})
	})
})
