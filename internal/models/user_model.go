package models

import (
	"time"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	StudentRole UserRole = "student"
	TeacherRole UserRole = "teacher"
)

type (
	User struct {
	ID         int      `gorm:"primaryKey"`
	Name       string   `gorm:"size:100;not null"`
	Email      string   `gorm:"not null;uniqueIndex"`
	Password   string   `gorm:"size:255;not null"`
	School     string   `gorm:"size:100;not null"`
	Role       UserRole `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

	Student struct {
		UserID         int `gorm:"primaryKey"`
		UserName       string
		UserSchool string
		Age            int
		Score          int
	}

	Teacher struct{
       UserID int `gorm:"primaryKey"`
       Age int
    }

	RefreshToken struct {
		ID           string
		UserID       int
		RefreshToken string
		ExpiresAt    time.Time
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)

