package models

import "time"

type SubmissionStatus string

const (
    StatusPending  SubmissionStatus = "pending"
    StatusApproved SubmissionStatus = "approved"
    StatusRejected SubmissionStatus = "rejected"
	StatusFailed SubmissionStatus = "failed"
)

type(
	Course struct{
		ID int `gorm:"primaryKey;autoIncrement;unique"`
		UserID int 
		Title string `gorm:"size:100;not null"`
		YoutubeUrl string `gorm:"size:500;not null"`
        YoutubeID  string `gorm:"size:50;not null;uniqueIndex"`
		CreatedAt time.Time
		User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	}

    CourseSubmission struct {
        ID int `gorm:"primaryKey;autoIncrement"`
		UserID int `gorm:"not null"`
		UserName string
		Title string `gorm:"size:200;not null"`
        YoutubeUrl string `gorm:"size:500;not null"`
        Status SubmissionStatus `gorm:"size:20;not null; default:'pending'"`
        RejectionReason string `gorm:"size:500"`
        CreatedAt time.Time
		UpdatedAt time.Time
        User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
    }
)