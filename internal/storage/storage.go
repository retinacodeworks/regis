package storage

import (
	"io"
	"path/filepath"
)

type storage interface {
	Upload(path string, archive io.Reader) error
	Has(path string) (bool, error)
	Get(path string) (any, error)
}

type Storage struct {
	backend storage
}

func New(backend storage) Storage {
	return Storage{backend}
}

func (s *Storage) UploadModule(path string, archive io.Reader) error {
	return s.backend.Upload(filepath.Join("modules", path), archive)
}

func (s *Storage) HasModule(path string) (bool, error) {
	return s.backend.Has(filepath.Join("modules", path))
}

func (s *Storage) GetModule(path string) (any, error) {
	return s.backend.Get(filepath.Join("modules", path))
}

func (s *Storage) UploadProvider(path string, archive io.Reader) error {
	return s.backend.Upload(filepath.Join("providers", path), archive)
}

func (s *Storage) HasProvider(path string) (bool, error) {
	return s.backend.Has(filepath.Join("providers", path))
}

func (s *Storage) GetProvider(path string) (any, error) {
	return s.backend.Get(filepath.Join("providers", path))
}
