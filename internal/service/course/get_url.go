package course

import (
	"context"
	"errors"
	"strings"
)
func(s *courseService) GetURL(ctx context.Context, objectKey string, baseUrl string) (string, error){

	if objectKey == "" {
		return "", errors.New("object key is empty")
	}


	url :=strings.TrimRight(baseUrl, "/") + "/" + strings.TrimLeft(objectKey, "/")
	return url, nil

}