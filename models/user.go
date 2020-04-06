package models

import (
	"time"
)

type Role string

type User struct {
	UserID          int    `json:"user_id"`
	Username        string    `json:"username"`
	Salt            string    `json:"salt"`
	Hash            string    `json:"hash"`
	IsAdministrator bool      `json:"is_administrator"`
	CreatedAt       time.Time `json:"created_at"`
}
