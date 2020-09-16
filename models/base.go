package models

import (
	"time"
)

type BaseModel struct {
	ID        uint64   `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
