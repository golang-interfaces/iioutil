package vioutil

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ VIOUtil

import (
	"io/ioutil"
	"os"
)

// virtual filesystem interface
type VIOUtil interface {
	// ReadDir reads the directory named by dirname and returns
	// a list of directory entries sorted by filename.
	ReadDir(dirname string) ([]os.FileInfo, error)

	// ReadFile reads the file named by filename and returns the contents.
	// A successful call returns err == nil, not err == EOF. Because ReadFile
	// reads the whole file, it does not treat an EOF from Read as an error
	// to be reported.
	ReadFile(filename string) ([]byte, error)

	// WriteFile writes data to a file named by filename.
	// If the file does not exist, WriteFile creates it with permissions perm;
	// otherwise WriteFile truncates it before writing.
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

func New() VIOUtil {
	return _VIOUtil{}
}

type _VIOUtil struct{}

func (this _VIOUtil) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func (this _VIOUtil) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (this _VIOUtil) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
