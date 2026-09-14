package course

import (
	"context"
	"errors"

	"studyhub/internal/models"
)

func (r *courseRepository) StoreFile(ctx context.Context, file models.File,) (int, error) {
	err := r.db.WithContext(ctx).Create(&file).Error
	if err != nil{
		return 0, errors.New("failed to save the main file")
	}

	return file.ID, nil
}