package course

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	dto "studyhub/internal/dto/course"
	"studyhub/internal/models"
)

func (s *courseService) StoreFile(ctx context.Context, file io.ReadSeeker, req dto.UploadFileRequest) (int, int, error) {

	if req.FileName == "" {
		return http.StatusBadRequest, 0, errors.New("file name is required")
	}


	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return http.StatusInternalServerError, 0, fmt.Errorf("calculate file hash: %w", err)
	}

	fileHash := hex.EncodeToString(hash.Sum(nil))

	// Reset file back to beginning.
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return http.StatusInternalServerError, 0, fmt.Errorf("reset file: %w", err)
	}

	

	extension := filepath.Ext(req.FileName)

	objectKey := fmt.Sprintf(
		"files/%s%s",
		fileHash,
		strings.ToLower(extension),
	)

	buffer := make([]byte, 512)

    n, err := file.Read(buffer)
    if err != nil {
        return http.StatusInternalServerError, 0 ,errors.New("failed to get mime type")
    }

mimeType := http.DetectContentType(buffer[:n])

_, err = file.Seek(0, io.SeekStart)

	if err := s.Upload(
		ctx,
		objectKey,
		file,
		"uploads",
	); err != nil {
		return http.StatusInternalServerError, 0, fmt.Errorf("upload file: %w", err)
	}
    
    fileModel := models.File{
		FileName: req.FileName,
		MimeType: mimeType,
		ObjectKey: objectKey,
		Hash: fileHash,
	}
    
	id, err := s.repo.StoreFile(ctx, fileModel)
	if err != nil {
    _ = os.Remove(filepath.Join("uploads", objectKey))
		return http.StatusInternalServerError, 0, fmt.Errorf("store file metadata: %w", err)
	}
    
	return http.StatusCreated, id, nil
}