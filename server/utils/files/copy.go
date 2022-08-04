package files

import (
	"github.com/pkg/errors"
	"io/ioutil"
)

func Copy(src, dest string) error {
	bytesRead, err := ioutil.ReadFile(src)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	err = ioutil.WriteFile(dest, bytesRead, 0644)
	if err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}