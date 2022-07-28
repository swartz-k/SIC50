package image

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func Save(fh *multipart.FileHeader, p string) error {
	d := path.Dir(p)
	if _, e := os.Open(d); os.IsNotExist(e) {
		os.Mkdir(d, 0777)
	}
	f, err := fh.Open()
	if err != nil {
		return errors.Wrap(err, "save file")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "read file for save")
	}
	return ioutil.WriteFile(p, b, 0644)
}
