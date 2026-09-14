package models

import "time"

type(
	Course struct{
		ID int `gorm:"primaryKey;autoIncrement;unique"`
		UserID int 
		Title string `gorm:"size:100;not null"`
		FileID int 
		CreatedAt time.Time
		File File `gorm:"foreignKey:FileID"`
		User User `gorm:"foreignKey:UserID"`
	}

	File struct {
        ID        int   `gorm:"primaryKey"`
        ObjectKey string `gorm:"size:500;not null;uniqueIndex"`
        FileName  string `gorm:"size:255"`
        MimeType  string `gorm:"size:100"`
        Hash      string `gorm:"size:64;uniqueIndex"`
    }

	CourseWithFile struct {
	    ID        int
		UserID int 
	    Title     string
	    ObjectKey string
		CreatedAt time.Time
    }
)