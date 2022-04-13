package utils

import (
	"io"
	"io/fs"
	"os"
)

// MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
func MkdirAll(path string) error {
	i := len(path)

	for i > 0 && !os.IsPathSeparator(path[i-1]) {
		i--
	}
	if i > 0 {
		err := os.MkdirAll(path[:i-1], os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create creates the named file with mode 0666 (before umask), truncating it if it already exists.
func Create(path string) (*os.File, error) {
	err := MkdirAll(path)
	if err != nil {
		return nil, err
	}
	return os.Create(path)
}

// OpenFile opens the named file with specified flag (O_RDONLY etc.).
func OpenFile(path string, flag int, perm fs.FileMode) (*os.File, error) {
	err := MkdirAll(path)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, flag, perm)
}

// CopyFile copies the src file to the dst path.
func CopyFile(dst string, src string) (int64, error) {
	s, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer s.Close()

	d, err := Create(dst)
	if err != nil {
		return 0, err
	}
	defer d.Close()

	return io.Copy(d, s)
}
