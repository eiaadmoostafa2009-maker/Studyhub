package models

import "time"
type Post struct {
	ID        int  `gorm:"primaryKey;autoIncrement;unique"`
	UserID    int  
	UserName string `gorm:"size:100;not null"`
	Title     string `gorm:"size:100;not null"`
	Content   string `gorm:"size:1000;not null"`
	Members []User `gorm:"many2many:post_members;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  
	UpdatedAt time.Time  
	IsJoined bool
}
