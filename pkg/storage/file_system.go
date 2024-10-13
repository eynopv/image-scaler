package storage

import (
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FileSystemStorage struct {
	directory string
}

func (s *FileSystemStorage) FilePath(fileName string) string {
	fp := filepath.Join(s.directory, fileName)
	return fp
}

func NewFileSystemStorage(directory string) *FileSystemStorage {
	return &FileSystemStorage{
		directory: directory,
	}
}

func (s *FileSystemStorage) Get(key string) (*os.File, error) {
	filepath := s.FilePath(key)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *FileSystemStorage) Save(key string, data io.Reader) (string, error) {
	if key == "" {
		key = uuid.New().String()
	}
	filepath := s.FilePath(key)

	dst, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, data)
	if err != nil {
		return "", err
	}

	return key, nil
}
