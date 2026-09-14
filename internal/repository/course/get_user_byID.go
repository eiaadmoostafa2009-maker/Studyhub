package course

import (
	"context"
	"studyhub/internal/models"

)

func (r *courseRepository) GetUserByID(ctx context.Context, userID int) (string, error) {
	//query to get user by id
	var user models.User
	User := r.db.First(&user, userID)
	
	if err :=User.Error; err != nil {
		return "", err
	}
	return user.Name, nil
}