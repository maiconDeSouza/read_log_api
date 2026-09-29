package models

import (
	"time"
	"uuid"
)

type BookStatus string

type User struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Nickname      string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"nickname"`
	Email         string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password      string     `gorm:"not null" json:"-"`
	IsVerified    bool       `gorm:"default:false" json:"isVerified"`
	CurrentBookID *uuid.UUID `gorm:"type:uuid" json:"currentBookID"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	UserBooks     []UserBook `gorm:"foreignKey:UserID" json:"userBooks"`
}

type UserBook struct {
	UserID       uuid.UUID  `gorm:"type:uuid;primaryKey" json:"userID"`
	BookID       uuid.UUID  `gorm:"type:uuid;primaryKey" json:"bookID"`
	CoverPath    string     `gorm:"type:varchar(255)"`
	Status       BookStatus `gorm:"type:varchar(20);not null" json:"status"`
	TotalPages   uint       `gorm:"not null" json:"totalPages"`
	PagesRead    uint       `gorm:"not null" json:"pagesRead"`
	SummaryPath  string     `gorm:"type:varchar(255)" json:"summaryPath,omitempty"`
	Observations string     `gorm:"type:text" json:"observations"`
	User         User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Book         Book       `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
}
