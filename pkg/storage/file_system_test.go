package storage

import "testing"

func TestNewFileSystemStorage(t *testing.T) {
	fs := NewFileSystemStorage("testdirectory")
	if fs.directory != "testdirectory" {
		t.Errorf("expected %v got %v; FileSystemStorage.directory", "testdirectory", fs.directory)
	}
}

func TestFileSystemStorageFilePath(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		expected string
	}{
		{"normal file", "image.jpg", "uploads/image.jpg"},
		{"file in subfolder", "folder/image.jpg", "uploads/folder/image.jpg"},
		{"file without extension", "folder/image", "uploads/folder/image"},
	}

	fs := NewFileSystemStorage("./uploads")
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			filePath := fs.FilePath(testCase.file)
			if filePath != testCase.expected {
				t.Errorf(
					"expected %v got %v; FileSystemStorage.FilePath(%v)",
					testCase.expected,
					filePath,
					testCase.file,
				)
			}
		})
	}
}
