package models
import "time"

type Post struct {
	ID        int  `gorm:"primaryKey;autoIncrement;unique"`
	UserID    int  
	Name string 
	Title     string `gorm:"size:100;not null"`
	Content   string `gorm:"size:1000;not null"`
	Members []User `gorm:"many2many:post_members;constraint:OnDelete:CASCADE"`
	MembersCount int `gorm:"column:members_count"`
	CreatedAt time.Time  
	UpdatedAt time.Time
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	IsJoined bool
}
