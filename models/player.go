package models

type Player struct {
	Model
	FirstName string `json:"first_name" binding:"required" validate:"min=1"`
	LastName string `json:"last_name" binding:"required" validate:"min=1"`
	Age int `json:"age" binding:"required" validate:"min=1"`
	//DateOfBirth time.Time `json:"date_of_birth" time_format:"2006-01-02" time_utc:"1"` FIXME couldn't get this working
}
