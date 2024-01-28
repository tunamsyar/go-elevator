package models

import (
	"time"
)

type Floor struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
