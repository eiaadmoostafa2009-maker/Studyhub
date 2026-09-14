package course

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

func (s *courseService) Upload(ctx context.Context, objectKey string,file io.Reader, baseUrl string) error {
	path := filepath.Join(baseUrl, objectKey)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	dst, err := os.Create(path)
	if err != nil {
		
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}