package storage

import "io"

type Azure struct {
}

func (a Azure) Upload(path string, archive io.Reader) error {
	return nil
}

func (a Azure) Has(path string) (bool, error) {
	return true, nil
}

func (a Azure) Get(path string) (io.Reader, error) {
	return nil, nil
}
