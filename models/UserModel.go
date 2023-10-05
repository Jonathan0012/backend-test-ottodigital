package models

type Users struct {
	Id         int64  `json:"id"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
