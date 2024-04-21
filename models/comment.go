package models

import "time"

type Comment struct {
	ID        uint      `json:"id" db:"id"`
	Text      string    `json:"text" db:"text"`
	UserId    uint      `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
