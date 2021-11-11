package models

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created"`
	UpdatedAt int64  `json:"updated"`
}
