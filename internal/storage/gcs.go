package storage

import "io"

type Gcs struct {
}

func (g Gcs) Upload(path string, archive io.Reader) error {
	return nil
}

func (g Gcs) Has(path string) (bool, error) {
	return true, nil
}

func (g Gcs) Get(path string) (io.Reader, error) {
	return nil, nil
}
