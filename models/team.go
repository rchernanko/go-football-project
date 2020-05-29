package models

type Team struct {
	Model
	Name string `json:"name" binding:"required" validate:"min=1"`
	Stadium uint `json:"stadium" binding:"required"`
}
