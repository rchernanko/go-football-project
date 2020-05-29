package models

import "time"

type Model struct {
	Id uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
