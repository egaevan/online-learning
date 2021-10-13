package model

type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Role     int    `json:"role"`
	Token    string `json:"token"`
}
