package models

type Tasks struct {
	Id          int64  `json:"id"`
	User_Id     int64  `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
