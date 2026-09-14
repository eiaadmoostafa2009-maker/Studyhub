package user

import "studyhub/internal/models"

type (
	//register
	RegisterRequest struct {
		UserName string         `json:"user_name" form:"name" validate:"required,min=3"`
		Email    string         `json:"email" form:"email" validate:"required,email"`
		Password string         `json:"password" form:"password" validate:"required"`
		School   string         `json:"university" form:"school" validate:"required"`
		UserRole models.UserRole `json:"user_role" form:"role" validate:"required,oneof=student teacher"`
	}

	//login
	LoginRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}
    
	RefreshTokenRequest struct{
		ID int
	}
)
