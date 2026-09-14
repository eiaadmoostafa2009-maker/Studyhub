package course

import (
	"context"

	"studyhub/internal/models"
)

func (r *courseRepository) GetFileByHash (ctx context.Context,	hash string,) (*models.File, error) {

	var file models.File

	err := r.db.WithContext(ctx).
		Where("hash = ?", hash).
		First(&file).Error

	if err != nil {
		return nil, err
	}

	return &file, nil
}