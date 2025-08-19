package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Author    string         `gorm:"not null" json:"author"`
	Desc      string         `json:"description"`
	UserID    uint           `json:"user_id"`
	User      *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
