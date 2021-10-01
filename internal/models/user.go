package models

type User struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
