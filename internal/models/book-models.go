package models

import "uuid"

type Book struct {
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title string    `gorm:"type:varchar(255);not null" json:"title"`
}
